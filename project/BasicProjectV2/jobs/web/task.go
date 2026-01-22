package web

import (
	"net/http"

	"github.com/basicprojectv2/jobs/domain"
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
	rg.POST("/add", r.HandleAddTask)
	rg.GET("/calc", r.ReCalcHotList)

}

func (r *TaskHandler) HandleAddTask(c *gin.Context) {
	req := domain.AddTaskReq{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := r.svc.AddTask(req, c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "task added"})
}

func (r *TaskHandler) ReCalcHotList(c *gin.Context) {
	if err := r.svc.ReCalcHotList(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "ok"})
}
