package service

import (
	"context"
	"github.com/robfig/cron/v3"
	"log"
	"timerDemo/dao"
	"timerDemo/model"
	"timerDemo/scheduler"
)

type TaskService interface {
	CreateTask(task model.TbTasks) error
	PauseTask(ctx context.Context, taskID uint) error
}

type taskService struct {
	repo      dao.TaskDAO
	scheduler *scheduler.CronScheduler
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

func NewTaskService(repo dao.TaskDAO, cronScheduler *scheduler.CronScheduler) TaskService {
	return taskService{repo: repo, scheduler: cronScheduler}
}

func (t taskService) PauseTask(ctx context.Context, taskID uint) (err error) {
	// 查询任务
	task, err := t.repo.GetTaskByID(ctx, taskID)
	if err != nil {
		return err
	}
	task.Status = 1 // 标记为暂停
	if err = t.repo.UpdateTask(task); err != nil {
		return err
	}
	if srErr := t.scheduler.RemoveTask(taskID); srErr != nil {
		return srErr
	}

	return nil
}
