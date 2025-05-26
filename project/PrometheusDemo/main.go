package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	app := InitializeApp()
	initPrometheus()
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
		log.Println("init Prometheus metrics server success")
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			log.Println("init Prometheus config error:", err)
			return
		}
	}()
}
