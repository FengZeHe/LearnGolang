package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type MonitoringBuilder struct {
	NameSpace string
	Subsystem string
	Name      string
	Help      string
}

func (m *MonitoringBuilder) HttpRequestTotalCounter() gin.HandlerFunc {
	httpRequestTotalCount := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: m.NameSpace,
			Subsystem: m.Subsystem,
			Name:      m.Name + "_Total_Count",
			Help:      m.Help + "_Total_Count_Help",
		},
		[]string{"path", "method"})
	prometheus.MustRegister(httpRequestTotalCount)
	return func(c *gin.Context) {
		httpRequestTotalCount.With(prometheus.Labels{"path": c.Request.URL.Path, "method": c.Request.Method}).Inc()
		c.Next()
	}
}

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
		duration := time.Since(start).Milliseconds()
		httpResponseTime.WithLabelValues(c.Request.URL.Path, c.Request.Method).Observe(float64(duration))
	}
}
