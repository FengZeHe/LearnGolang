package dao

import (
	"context"
	"gorm.io/gorm"
	"log"
	"timerDemo/model"
)

type TaskDAO interface {
	InsertTask(task model.TbTasks) error
	GetActiveTasks() ([]*model.TbTasks, error)
	UpdateTask(task *model.TbTasks) error
	GetTaskByID(ctx context.Context, id uint) (*model.TbTasks, error)
}
type GromTaskDAO struct {
	DB *gorm.DB
}

func NewTaskDAO(db *gorm.DB) TaskDAO {
	return &GromTaskDAO{DB: db}
}

func (dao GromTaskDAO) InsertTask(task model.TbTasks) (err error) {
	result := dao.DB.Table("tasks").Create(&task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 获取正在执行的任务
func (dao GromTaskDAO) GetActiveTasks() (tasks []*model.TbTasks, err error) {
	res := dao.DB.Table("tasks").Where("status = ?", 0).Find(&tasks)
	if res.Error != nil {
		log.Println(res.Error)
		return nil, res.Error
	}
	return tasks, nil
}

func (dao GromTaskDAO) GetTaskByID(ctx context.Context, id uint) (*model.TbTasks, error) {
	var task model.TbTasks
	if err := dao.DB.Table("tasks").WithContext(ctx).First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (dao GromTaskDAO) UpdateTask(task *model.TbTasks) (err error) {
	result := dao.DB.Table("tasks").Save(&task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
