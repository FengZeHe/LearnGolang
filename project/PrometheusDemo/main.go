package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	app := InitializeApp()

	initPrometheus()

	server := app.server
	server.Use(func(c *gin.Context) {
		httpRequestsTotal.WithLabelValues(c.Request.URL.Path, c.Request.Method).Inc()
		c.Next()
	})
	err := server.Run(":8088")
	if err != nil {
		return
	}
}

func initPrometheus() {
	go func() {
		// 专门给 prometheus 用的端口
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8081", nil)
	}()
}
