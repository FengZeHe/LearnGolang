package dao

import (
	"context"
	"log"
	"time"

	"github.com/basicprojectv2/jobs/domain"
	"gorm.io/gorm"
)

type GromTbTask struct {
	db *gorm.DB
}

type TaskDAO interface {
	AddTask(req domain.Task, ctx context.Context) (err error)
	UpdateTask(req domain.Task) (err error)
	DeleteTask(req domain.DeleteTaskReq) (err error)
	FindTaskByID(id string, ctx context.Context) (task domain.Task, err error)
	FindAllTasks(req domain.TaskFilterReq, ctx context.Context) (tasks []domain.Task, err error)
	FindActiveTasks() (tasks []domain.Task, err error)
}

func NewTaskDAO(db *gorm.DB) TaskDAO {
	return &GromTbTask{db: db}
}

func (t *GromTbTask) AddTask(req domain.Task, ctx context.Context) (err error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	req.CreatedAt = now
	res := t.db.WithContext(ctx).Create(&req)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (t *GromTbTask) UpdateTask(req domain.Task) (err error) {
	res := t.db.Save(&req)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (t *GromTbTask) DeleteTask(req domain.DeleteTaskReq) (err error) {
	res := t.db.Delete(&req)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (t *GromTbTask) FindTaskByID(id string, ctx context.Context) (task domain.Task, err error) {
	res := t.db.WithContext(ctx).Where("id = ?", id).First(&task)
	if res.Error != nil {
		return task, res.Error
	}
	return task, nil
}

func (t *GromTbTask) FindAllTasks(req domain.TaskFilterReq, ctx context.Context) (tasks []domain.Task, err error) {
	if req.Status > 0 {
		res := t.db.WithContext(ctx).Find(&tasks).Where("status = ?", req.Status)
		if res.Error != nil {
			return tasks, res.Error
		}
		log.Println(res)
		return tasks, nil
	}

	res := t.db.WithContext(ctx).Find(&tasks)
	if res.Error != nil {
		return tasks, res.Error
	}
	return tasks, nil
}

func (t *GromTbTask) FindActiveTasks() (tasks []domain.Task, err error) {
	res := t.db.Where("status = ?", 0).Find(&tasks)
	if res.Error != nil {
		return tasks, res.Error
	}
	return tasks, nil

}
