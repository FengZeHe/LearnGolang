package service

import (
	"context"

	"github.com/basicprojectv2/jobs/domain"
	"github.com/basicprojectv2/jobs/repository"
	"github.com/basicprojectv2/jobs/scheduler"
)

type taskService struct {
	taskRepo  repository.TaskRepository
	scheduler *scheduler.CronScheduler
}

type TaskService interface {
	AddTask(req domain.AddTaskReq, ctx context.Context) (err error)
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

func (t taskService) StartTask() {

}
