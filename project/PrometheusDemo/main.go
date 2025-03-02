package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	app := InitializeApp()

	initPrometheus()

	app.cron.Start()
	defer func() {
		// 等待定时任务退出
		<-app.cron.Stop().Done()
	}()

	server := app.server
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
