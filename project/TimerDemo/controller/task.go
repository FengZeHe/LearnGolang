package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.JSON(http.StatusOK, gin.H{"data": "okk"})

}
