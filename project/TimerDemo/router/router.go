package router

import (
	"github.com/gin-gonic/gin"
	"timerDemo/controller"
	"timerDemo/service"
)

// 注册路由
func RegisterTaskRouter(srv service.TaskService) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	tc := controller.NewTaskController(srv)
	v1 := r.Group("/task")
	{
		// 任务部分
		v1.POST("/add", tc.AddTask)
		v1.POST("/update")
		v1.POST("/delete/:id")
		v1.GET("/")
		v1.POST("/stop/:id", tc.PauseTask)
		// 日志部分
		v1.GET("/log")
		v1.GET("/log/:id")
	}
	return r
}
