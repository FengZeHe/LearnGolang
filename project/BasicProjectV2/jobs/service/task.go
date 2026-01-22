package service

import (
	"context"

	"github.com/basicprojectv2/jobs/domain"
	"github.com/basicprojectv2/jobs/repository"
	"github.com/basicprojectv2/jobs/scheduler"
	"github.com/gin-gonic/gin"
)

type taskService struct {
	taskRepo  repository.TaskRepository
	scheduler *scheduler.CronScheduler
}

type TaskService interface {
	AddTask(req domain.AddTaskReq, ctx context.Context) (err error)
	ReCalcHotList(c *gin.Context) (err error)
}

func NewTaskService(taskRepo repository.TaskRepository, scheduler *scheduler.CronScheduler) TaskService {
	return &taskService{
		taskRepo:  taskRepo,
		scheduler: scheduler,
	}
}

func (t taskService) AddTask(req domain.AddTaskReq, ctx context.Context) (err error) {
	return t.taskRepo.AddTask(req, ctx)
}

func (t taskService) ReCalcHotList(c *gin.Context) (err error) {
	return t.taskRepo.ReCalcHotList(c)
}

func (t taskService) StartTask() {

}
