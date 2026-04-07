package service

import (
	"context"
	"log"

	"github.com/basicprojectv2/jobs/code"
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
	UpdateTask(req domain.UpdateTaskReq, ctx context.Context) (err error)
	DeleteTask(req domain.DeleteTaskReq, ctx context.Context) (err error)
	ReCalcHotList(c *gin.Context) (err error)
	GetTasksList(req domain.PageReq, ctx context.Context) (d domain.PageResp, err error)
	GetTask(req domain.TaskReq, ctx context.Context) (d domain.Task, err error)
	StartTask(req domain.TaskReq, ctx context.Context) (err error)
	PauseTask(req domain.TaskReq, ctx context.Context) (err error)
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

func (t taskService) UpdateTask(req domain.UpdateTaskReq, ctx context.Context) (err error) {
	if err = t.scheduler.RemoveTask(req.ID); err != nil {
		return err
	}
	return nil
}

func (t taskService) DeleteTask(req domain.DeleteTaskReq, ctx context.Context) (err error) {
	if err = t.scheduler.RemoveTask(req.ID); err != nil {
		return err
	}
	return t.taskRepo.DeleteTask(req, ctx)
}

func (t taskService) StartTask(req domain.TaskReq, ctx context.Context) (err error) {
	task, err := t.taskRepo.QueryTaskByID(req, ctx)
	if err != nil {
		return err
	}
	if err = t.scheduler.StartTask(task); err != nil {
		return err
	}
	if err = t.taskRepo.UpdateTaskStatus(req, ctx, code.TaskRunning); err != nil {
		return err
	}
	return nil
}

func (t taskService) PauseTask(req domain.TaskReq, ctx context.Context) (err error) {
	task, err := t.taskRepo.QueryTaskByID(req, ctx)
	if err != nil {
		return err
	}

	if err = t.scheduler.RemoveTask(task.ID); err != nil {
		return err
	}
	if err = t.taskRepo.UpdateTaskStatus(req, ctx, code.TaskPause); err != nil {
		return err
	}

	return nil

}

func (t taskService) GetTasksList(req domain.PageReq, ctx context.Context) (d domain.PageResp, err error) {
	d, err = t.taskRepo.GetAllTasks(req, ctx)
	if err != nil {
		return d, err
	}
	return d, nil
}

func (t taskService) GetTask(req domain.TaskReq, ctx context.Context) (d domain.Task, err error) {
	d, err = t.taskRepo.QueryTaskByID(req, ctx)
	if err != nil {
		return d, err
	}
	log.Print(d)
	return d, nil
}

func (t taskService) ReCalcHotList(c *gin.Context) (err error) {
	return t.taskRepo.ReCalcHotList(c)
}
