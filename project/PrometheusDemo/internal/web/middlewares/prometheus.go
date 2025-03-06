package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"log"
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
			Buckets:   prometheus.ExponentialBuckets(0.01, 2, 10),
		},
		[]string{"path", "method"})
	prometheus.MustRegister(httpRequestDuration)
	return func(c *gin.Context) {
		start := time.Now()
		time.Sleep(200 * time.Millisecond)
		c.Next()
		duration := time.Since(start).Milliseconds()
		log.Println("duration=", duration)
		httpRequestDuration.WithLabelValues(c.Request.URL.Path, c.Request.Method).Observe(float64(duration))
	}
}
