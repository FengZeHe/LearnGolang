package ioc

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

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
