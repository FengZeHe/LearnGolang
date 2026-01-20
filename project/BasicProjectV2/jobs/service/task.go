package service

import (
	"context"

	"github.com/basicprojectv2/jobs/domain"
	"github.com/basicprojectv2/jobs/repository"
)

type taskService struct {
	taskRepo repository.TaskRepository
}

type TaskService interface {
	AddTask(req domain.AddTask, ctx context.Context) (err error)
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (t taskService) AddTask(req domain.AddTask, ctx context.Context) (err error) {
	return t.taskRepo.AddTask(req, ctx)
}
