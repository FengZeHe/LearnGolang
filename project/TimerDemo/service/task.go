package service

import (
	"github.com/robfig/cron/v3"
	"log"
	"timerDemo/dao"
	"timerDemo/model"
)

type TaskService interface {
	CreateTask(task model.TbTasks) error
}

type taskService struct {
	repo dao.TaskDAO
}

func (t taskService) CreateTask(task model.TbTasks) (err error) {

	parser := cron.NewParser(
		cron.SecondOptional | // 允许秒字段（可选）
			cron.Minute | // 分钟字段
			cron.Hour | // 小时字段
			cron.Dom | // 日期字段
			cron.Month | // 月份字段
			cron.Dow | // 星期字段
			cron.Descriptor, // 支持描述符，如@daily, @weekly等
	)

	_, err = parser.Parse(task.CronExpr)
	if err != nil {
		log.Println("cron expr error: %v\n", err)
		return err
	}

	task.Status = 0
	if err = t.repo.InsertTask(task); err != nil {
		return err
	}
	return nil
}

func NewTaskService(repo dao.TaskDAO) TaskService {
	return taskService{repo: repo}
}
