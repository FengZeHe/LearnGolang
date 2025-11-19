package ioc

import (
	"time"

	intWeb "github.com/basicprojectv2/interactive/web"
	"github.com/basicprojectv2/internal/web"
	"github.com/basicprojectv2/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// 初始化gin Engine
func InitWebServer(mdls []gin.HandlerFunc, userHdl *web.UserHandler, sysHdl *web.SysHandler,
	menuHdl *web.MenuHandler, roleHdl *web.RoleHandler, draftHdl *web.DraftHandler,
	articleHdl *web.ArticleHandler, commentHdl *web.CommentHandler, intHdl *intWeb.InteractiveHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls[0], mdls[4], mdls[5], mdls[6])
	userHdl.RegisterRoutes(server, mdls[3], mdls[2])
	sysHdl.RegisterRoutes(server, mdls[1], mdls[2], mdls[3])
	menuHdl.RegisterRoutes(server, mdls[2])
	roleHdl.RegisterRoutes(server, mdls[2])
	draftHdl.RegisterRoutes(server, mdls[2])
	articleHdl.RegisterRoutes(server, mdls[2])
	commentHdl.RegisterRoutes(server, mdls[2])
	intHdl.RegisterRoutes(server, mdls[2])

	return server
}

// 初始化中间件
func InitGinMiddlewares(ca *middleware.CasbinRoleCheck, i18n *i18n.Bundle) []gin.HandlerFunc {
	pb := &middleware.MonitoringBuilder{
		NameSpace: "basic_projectV2",
		Subsystem: "ordinary_service",
		Name:      "http_request",
		Help:      "",
	}

	return []gin.HandlerFunc{
		cors.New(cors.Config{
			//AllowAllOrigins: true,
			//AllowOrigins:     []string{"http://localhost:3000"},
			AllowCredentials: true,

			AllowHeaders: []string{"Content-Type", "Authorization"},
			// 这个是允许前端访问你的后端响应中带的头部
			ExposeHeaders: []string{"x-jwt-token"},
			//AllowHeaders: []string{"content-type"},
			//AllowMethods: []string{"POST"},
			//AllowOriginFunc: fun(origin string) bool {
			//	if strings.HasPrefix(origin, "http://localhost") {
			//		//if strings.Contains(origin, "localhost") {
			//		return true
			//	}
			//	return strings.Contains(origin, "your_company.com")
			//},
			AllowAllOrigins: true,
			MaxAge:          12 * time.Hour,
		}),
		ca.CheckRole(),
		(&middleware.LoginJWTMiddlewareBuilder{}).CheckLogin(),
		middleware.I18nMiddleware(i18n),
		pb.HttpRequestTotalCounter(),
		pb.HttpResponseTime(),
		otelgin.Middleware("gin-service"),
	}
}
