package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type App struct {
	server *gin.Engine
	cron   *cron.Cron
}

// 定义一个计数器
//var httpRequestsTotal = prometheus.NewCounterVec(
//	prometheus.CounterOpts{
//		Name: "http_requests_total",
//		Help: "Total number of HTTP requests.",
//	},
//	[]string{"path", "method"},
//)

//func init() {
//	// 注册计数器
//	prometheus.MustRegister(httpRequestsTotal)
//}
