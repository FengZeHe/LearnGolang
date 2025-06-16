package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"timerDemo/model"
	"timerDemo/service"
)

type TaskController struct {
	srv service.TaskService
}

func NewTaskController(srv service.TaskService) *TaskController {
	return &TaskController{srv: srv}
}

func (tc *TaskController) AddTask(c *gin.Context) {
	req := model.AddTaskReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := model.TbTasks{
		Name:     req.Name,
		CronExpr: req.CronExpr,
		TaskType: req.TaskType,
		TaskName: req.TaskName,
	}
	if err := tc.srv.CreateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "task added"})
}

func (tc *TaskController) PauseTask(c *gin.Context) {
	taskIDStr := c.Param("id")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 调用service层停止服务
	ctx := context.Background()
	if paErr := tc.srv.PauseTask(ctx, uint(taskID)); paErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": paErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "任务已暂停"})
}
