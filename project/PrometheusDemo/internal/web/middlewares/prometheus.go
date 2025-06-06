package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"net"
	"time"
)

type MonitoringBuilder struct {
	NameSpace string
	Subsystem string
	Name      string
	Help      string
}

// 也可以这样实现
func NewMonitoringBuilder(namespace, subsystem, name, help string) *MonitoringBuilder {
	return &MonitoringBuilder{
		NameSpace: namespace,
		Subsystem: subsystem,
		Name:      name,
		Help:      help,
	}
}

func (m *MonitoringBuilder) HttpRequestTotalCounter() gin.HandlerFunc {
	httpRequestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: m.NameSpace,
			Subsystem: m.Subsystem,
			Name:      m.Name + "_Total_Count",
			Help:      m.Help,
		},
		[]string{"path", "method"},
	)
	prometheus.MustRegister(httpRequestsTotal)
	return func(c *gin.Context) {
		httpRequestsTotal.With(prometheus.Labels{"path": c.Request.URL.Path, "method": c.Request.Method}).Inc()
		c.Next()
	}
}

func (m *MonitoringBuilder) HttpRequestDurationHistogram() gin.HandlerFunc {
	httpRequestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: m.NameSpace,
			Subsystem: m.Subsystem,
			Name:      m.Name + "_Duration",
			Help:      m.Help,
			//Buckets:   prometheus.ExponentialBuckets(0.01, 2, 10),
			Buckets: []float64{100, 200, 300, 400, 500},
		},
		[]string{"path", "method"})
	prometheus.MustRegister(httpRequestDuration)
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()
		//log.Println("duration=", duration)
		httpRequestDuration.WithLabelValues(c.Request.URL.Path, c.Request.Method).Observe(float64(duration))
	}
}

// summary指标记录HTTP请求响应时间
func (m *MonitoringBuilder) HttpResponseTime() gin.HandlerFunc {
	httpResponseTime := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  m.NameSpace,
			Subsystem:  m.Subsystem,
			Name:       m.Name + "_ResponseTime",
			Help:       m.Help,
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"path", "method"},
	)
	prometheus.MustRegister(httpResponseTime)
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()
		httpResponseTime.WithLabelValues(c.Request.URL.Path, c.Request.Method).Observe(duration)
	}
}

var gormQueryDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: "my_gorm",
		Subsystem: "prometheus_demo",
		Name:      "gorm_query_duration_seconds",
		Help:      "xxx",
		Buckets:   []float64{0.001, 0.01, 0.1, 1, 10},
	},
	[]string{"query_type"},
)

func init() {
	prometheus.MustRegister(gormQueryDuration)
}

// GORM回调函数，用于统计查询时间
func GormQueryCallback(db *gorm.DB) {
	if db.Error != nil {
		return
	}

	// 在查询执行前记录开始时间
	db.Callback().Query().Before("gorm:query").Register("prometheus:record_start", func(db *gorm.DB) {
		db.InstanceSet("query_start_time", time.Now())
	})

	// 等待查询完成
	db.Callback().Query().After("gorm:query").Register("prometheus:query_duration", func(db *gorm.DB) {
		if start, ok := db.InstanceGet("query_start_time"); ok {
			if startTime, ok := start.(time.Time); ok {
				duration := time.Since(startTime).Seconds()
				queryType := "unknown"
				if db.Statement.Schema != nil {
					queryType = db.Statement.Schema.Table
				}
				gormQueryDuration.WithLabelValues(queryType).Observe(duration)
				log.Println(queryType, db.Statement.Schema.Name, duration)
			}
		}
	})
}

type CacheHitHook struct {
	NameSpace        string
	Subsystem        string
	Name             string
	Help             string
	cacheHitCounter  *prometheus.CounterVec
	cacheMissCounter *prometheus.CounterVec
}

func (h *CacheHitHook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	}
}

func (h *CacheHitHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	cacheHitCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: h.NameSpace,
			Subsystem: h.Subsystem,
			Name:      h.Name + "+_CacheHit",
			Help:      h.Help,
		},
		[]string{"operation"})

	cacheMissCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: h.NameSpace,
			Subsystem: h.Subsystem,
			Name:      h.Name + "+_CacheMiss",
			Help:      h.Help,
		},
		[]string{"operation"})

	prometheus.MustRegister(cacheHitCounter, cacheMissCounter)
	return func(ctx context.Context, cmd redis.Cmder) (err error) {
		if cmd.Name() == "get" {
			err = next(ctx, cmd)
			if err == nil {
				// 命令执行成功，缓存命中
				log.Println("cache hit")
				cacheHitCounter.WithLabelValues("GET").Inc()
			} else if err == redis.Nil {
				// 键不存在，缓存未命中
				log.Println("cache miss")
				cacheMissCounter.WithLabelValues("GET").Inc()
			} else {
				// 其他错误
				log.Printf("Error executing GET command: %v", err)
			}

		}
		return nil
	}
}

func (h *CacheHitHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		return next(ctx, cmds)
	}
}

func NewCacheHitHook(namespace, subsystem, name string, help string) *CacheHitHook {
	return &CacheHitHook{
		NameSpace: namespace,
		Subsystem: subsystem,
		Name:      name,
		Help:      help,
	}
}

// 定义缓存命中率的计数器
func (h *CacheHitHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	cacheHitCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: h.NameSpace,
			Subsystem: h.Subsystem,
			Name:      h.Name + "+_CacheHit",
			Help:      h.Help,
		},
		[]string{"operation"})

	cacheMissCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: h.NameSpace,
			Subsystem: h.Subsystem,
			Name:      h.Name + "+_CacheMiss",
			Help:      h.Help,
		},
		[]string{"operation"})

	prometheus.MustRegister(cacheHitCounter, cacheMissCounter)

	if cmd.Name() == "GET" {
		if cmd.Err() == redis.Nil {
			// 未命中缓存
			cacheMissCounter.WithLabelValues("GET").Inc()
		} else if cmd.Err() == redis.Nil {
			// 命中缓存
			cacheHitCounter.WithLabelValues("GET").Inc()
		}
	}
	return nil
}

// 实现 Hook 接口的空方法
func (h *CacheHitHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	return ctx, nil
}

//fun (h *CacheHitHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
//	if cmd.Name() == "get" { // Redis 命令名称是小写的
//		if cmd.Err() == redis.Nil {
//			h.cacheMissCounter.WithLabelValues("get").Inc()
//		} else if cmd.Err() == nil {
//			h.cacheHitCounter.WithLabelValues("get").Inc()
//		}
//	}
//	return nil
//}

func (h *CacheHitHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (h *CacheHitHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	return nil
}
