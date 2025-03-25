package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0" // 使用统一的架构版本
)

func main() {
	// 初始化 Zipkin exporter 和 TracerProvider
	zipkinEndpoint := "http://localhost:9411/api/v2/spans"
	tp, err := initTracerProvider(zipkinEndpoint)
	if err != nil {
		log.Fatalf("failed to initialize tracer provider: %v", err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	r := gin.Default()
	r.Use(otelgin.Middleware("gin-service"))

	// 定义一个带有参数的 GET 路由
	r.GET("/hello", HandleHi)

	// 启动服务器，监听 8080 端口
	r.Run(":8080")
}

func HandleHi(c *gin.Context) {
	// 创建新的 span
	_, span := otel.Tracer("gin-service").Start(c.Request.Context(), "handleHello")
	defer span.End()

	// 添加 span 标签
	span.SetAttributes(attribute.String("http.path", c.Request.URL.Path))

	time.Sleep(time.Duration(rand.Intn(200)+1) * time.Millisecond)

	c.JSON(200, gin.H{
		"message": "hi!",
	})

	// 设置 span 状态
	span.SetStatus(codes.Ok, "Success")
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
