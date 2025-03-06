package ioc

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"prometheusdemo/internal/web"
	"prometheusdemo/internal/web/middlewares"
	"time"
)

func InitWebServer(mdls []gin.HandlerFunc, userHdl *web.UserHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	userHdl.RegisterRoutes(server)

	return server
}

func InitGinMiddlewares() []gin.HandlerFunc {
	pb := &middlewares.MonitoringBuilder{
		NameSpace: "my_http",
		Subsystem: "prometheus_demo",
		Name:      "my_http_request",
		Help:      "...",
	}
	//pb := middlewares.NewMonitoringBuilder("my_http_v2", "prometheus_demo", "my_http_request_total_count", "Total number of HTTP requests")

	return []gin.HandlerFunc{
		cors.New(cors.Config{
			AllowCredentials: true,
			AllowHeaders:     []string{"Content-Type", "Authorization"},
			ExposeHeaders:    []string{"x-jwt-token"},
			AllowAllOrigins:  true,
			MaxAge:           12 * time.Hour,
		}),
		pb.HttpRequestTotalCounter(),
		pb.HttpRequestDurationHistogram(),
	}
}
