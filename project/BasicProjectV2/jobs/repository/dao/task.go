package dao

import (
	"context"

	"github.com/basicprojectv2/jobs/domain"
	"gorm.io/gorm"
)

type GromTbTask struct {
	db *gorm.DB
}

type TaskDAO interface {
	AddTask(req domain.AddTask, ctx context.Context) (err error)
}

func NewTaskDAO(db *gorm.DB) TaskDAO {
	return &GromTbTask{db: db}
}

func (t *GromTbTask) AddTask(req domain.AddTask, ctx context.Context) (err error) {
	return nil
}
