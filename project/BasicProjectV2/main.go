package main

import (
	"context"

	"github.com/basicprojectv2/jobs"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"

	"log"

	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0" // 使用统一的架构版本
)

func main() {
	app := InitializeApp()
	p := initPrometheus()
	go func() {
		if err := p.Run(":8081"); err != nil {
			log.Fatal(err)
		}
	}()
	// 初始化定时任务注册
	initTaskRegistry(app.Registry)
	go func() {
		if err := app.Scheduler.Start(context.Background()); err != nil {
			log.Fatalln("启动 cron 失败:", err)
		}
	}()

	// 初始化 Zipkin exporter 和 TracerProvider
	//zipkinEndpoint := "http://192.168.95.131:9411/api/v2/spans"
	//tp, err := initTracerProvider(zipkinEndpoint)
	//if err != nil {
	//	log.Fatalf("failed to initialize tracer provider: %v", err)
	//}
	//defer func() {
	//	if err := tp.Shutdown(context.Background()); err != nil {
	//		log.Printf("Error shutting down tracer provider: %v", err)
	//	}
	//}()

	server := app.Server
	err := server.Run(":8088")
	if err != nil {
		return
	}

}

// 初始化Prometheus
func initPrometheus() *gin.Engine {

	var appAlive = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "app_alive",
		Help: "Heartbeat detection",
	})
	prometheus.MustRegister(appAlive)
	appAlive.Set(1)

	r := gin.Default()
	r.GET("/metrics", promAuthMiddleware(), gin.WrapH(promhttp.Handler()))
	log.Println("init Prometheus metrics Server success!")
	return r
}

func promAuthMiddleware() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"prometheus": "prometheus_password",
	})
}

func initTracerProvider(zipkinEndpoint string) (*tracesdk.TracerProvider, error) {
	// 创建 Zipkin exporter
	exp, err := zipkin.New(zipkinEndpoint)
	if err != nil {
		return nil, err
	}

	// 创建资源，添加服务名称等信息
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("gin-service"),
		),
	)
	if err != nil {
		return nil, err
	}

	// 创建 TracerProvider
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(r),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

func initTaskRegistry(tr *jobs.TaskRegistry) {
	//tr.Register("sayhi", events.ExecTimeKeeping)
}
