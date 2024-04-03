package ioc

import (
	"github.com/basicprojectv2/internal/web"
	"github.com/basicprojectv2/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// 初始化gin Engine
func InitWebServer(mdls []gin.HandlerFunc, userHdl *web.UserHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	// userHandler 注册user的路由
	userHdl.RegisterRoutes(server)
	return server
}

// 初始化中间件
func InitGinMiddlewares() []gin.HandlerFunc {
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
			//AllowOriginFunc: func(origin string) bool {
			//	if strings.HasPrefix(origin, "http://localhost") {
			//		//if strings.Contains(origin, "localhost") {
			//		return true
			//	}
			//	return strings.Contains(origin, "your_company.com")
			//},
			AllowAllOrigins: true,
			MaxAge:          12 * time.Hour,
		}),
		(&middleware.LoginJWTMiddlewareBuilder{}).CheckLogin(), // 什么写法？
	}
}

type UserHandler struct {
}
