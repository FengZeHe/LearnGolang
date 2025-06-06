package service

import (
	"timerDemo/dao"
	"timerDemo/model"
)

type TaskService interface {
	CreateTask(task model.TbTasks) error
}

type taskService struct {
	repo dao.TaskDAO
}

func (t taskService) CreateTask(task model.TbTasks) error {
	return nil
}

func NewTaskService(repo dao.TaskDAO) TaskService {
	return taskService{repo: repo}
}
