package repository

import (
	"context"

	"github.com/basicprojectv2/jobs/domain"
	"github.com/basicprojectv2/jobs/repository/dao"
)

type taskRepository struct {
	taskDAO dao.TaskDAO
}
type TaskRepository interface {
	AddTask(req domain.AddTask, ctx context.Context) (err error)
}

func NewTaskRepository(taskDAO dao.TaskDAO) TaskRepository {
	return &taskRepository{taskDAO: taskDAO}
}

func (t taskRepository) AddTask(req domain.AddTask, ctx context.Context) (err error) {
	return t.taskDAO.AddTask(req, ctx)
}
