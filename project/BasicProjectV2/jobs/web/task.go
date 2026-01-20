package web

import (
	"github.com/basicprojectv2/jobs/service"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	svc service.TaskService
}

func NewTaskHandler(svc service.TaskService) *TaskHandler {
	return &TaskHandler{svc: svc}
}

func (r *TaskHandler) RegisterRoutes(server *gin.Engine, loginCheck gin.HandlerFunc) {
	rg := server.Group("/v2/task/")
	rg.Use(loginCheck)
	rg.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	/*
		todo
		1. 获取所有定时任务
		2. 添加任务
		3. 删除任务
		4. 启动/暂停任务
		5. 查看某个定时任务
	*/
	rg.GET("/")

}
