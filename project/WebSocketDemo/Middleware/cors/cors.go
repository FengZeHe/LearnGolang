package Cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	c := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH"},         //允许的方法
		AllowHeaders:    []string{"Content-Type", "Access-Token", "Authorization"}, //允许设置的头部
		MaxAge:          6 * time.Hour,                                             // 设置过期时间
	}
	return cors.New(c)
}
