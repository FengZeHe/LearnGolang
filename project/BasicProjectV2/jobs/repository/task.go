package repository

import (
	"context"
	"log"

	"github.com/basicprojectv2/jobs/domain"
	"github.com/basicprojectv2/jobs/repository/dao"
	"github.com/robfig/cron/v3"
)

type taskRepository struct {
	taskDAO dao.TaskDAO
}
type TaskRepository interface {
	AddTask(req domain.AddTaskReq, ctx context.Context) (err error)
}

func NewTaskRepository(taskDAO dao.TaskDAO) TaskRepository {
	return &taskRepository{taskDAO: taskDAO}
}

func (t taskRepository) AddTask(req domain.AddTaskReq, ctx context.Context) (err error) {

	parser := cron.NewParser(
		cron.SecondOptional | // 允许秒字段（可选）
			cron.Minute | // 分钟字段
			cron.Hour | // 小时字段
			cron.Dom | // 日期字段
			cron.Month | // 月份字段
			cron.Dow | // 星期字段
			cron.Descriptor, // 支持描述符，如@daily, @weekly等
	)

	task := domain.Task{
		Name:     req.Name,
		TaskName: req.TaskName,
		TaskType: req.TaskType,
		CronExpr: req.CronExpr,
	}

	_, err = parser.Parse(task.CronExpr)
	if err != nil {
		log.Println("cron expr error: %v\n", err)
		return err
	}

	return t.taskDAO.AddTask(task, ctx)
}
