package dao

import (
	"gorm.io/gorm"
	"timerDemo/model"
)

type TaskDAO interface {
	InsertTask(task model.TbTasks) error
}
type GromTaskDAO struct {
	DB *gorm.DB
}

func NewTaskDAO(db *gorm.DB) TaskDAO {
	return &GromTaskDAO{DB: db}
}

func (dao GromTaskDAO) InsertTask(task model.TbTasks) error {
	// todo 检查cron表达式
	return nil
}
