package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建默认的 Gin 引擎，包含日志和恢复中间件
	r := gin.Default()

	// 定义一个简单的 GET 路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// 定义一个带有参数的 GET 路由
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name + "!",
		})
	})

	// 启动服务器，监听 8080 端口
	r.Run(":8080")
}
