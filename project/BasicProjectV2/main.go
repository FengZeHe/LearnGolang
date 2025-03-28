package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
)

func main() {
	app := InitializeApp()
	p := initPrometheus()
	go func() {
		if err := p.Run(":8081"); err != nil {
			log.Fatal(err)
		}
	}()

	server := app.server
	err := server.Run(":8088")
	if err != nil {
		return
	}

}

// 初始化Prometheus
func initPrometheus() *gin.Engine {
	r := gin.Default()
	r.GET("/metrics", promAuthMiddleware(), gin.WrapH(promhttp.Handler()))
	log.Println("init Prometheus metrics server success!")
	return r
}

func promAuthMiddleware() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"prometheus": "prometheus_password",
	})
}
