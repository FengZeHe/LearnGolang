package repository

import (
	"context"
	"log"

	"github.com/basicprojectv2/jobs/domain"
	"github.com/basicprojectv2/jobs/repository/dao"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
)

type taskRepository struct {
	taskDAO dao.TaskDAO
	rdb     redis.Cmdable
}
type TaskRepository interface {
	AddTask(req domain.AddTaskReq, ctx context.Context) (err error)
	ReCalcHotList(ctx context.Context) (err error)
}

func NewTaskRepository(taskDAO dao.TaskDAO, rdb redis.Cmdable) TaskRepository {
	return &taskRepository{taskDAO: taskDAO, rdb: rdb}
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

func (t taskRepository) ReCalcHotList(ctx context.Context) (err error) {
	/*
		todo 1. 分批查询. 每次1000篇，设置门槛：只查阅读量大于xx的，日榜只查今天的
	*/
	pageSize := 1000
	index := 1
	for {
		alist, err := t.taskDAO.GetArticleIDs(index, pageSize)
		if err != nil {
			log.Println("err", err)
			return err
		}
		if len(alist) == 0 {
			break
		}

		// todo 每计算完1000篇就写入redis,清空缓存区
		for _, item := range alist {
			log.Println(item.ID, item.Title)

			_, err := t.rdb.ZAdd(ctx, "hotlist/articles/score/", redis.Z{1, item.ID}).Result()
			if err != nil {
				log.Println("err", err)
				return
			}

			_, err = t.rdb.HSet(ctx, "hotlist/articles/"+item.ID, "title", item.Title, "score", 0.1).Result()
			if err != nil {
				log.Println("err", err)
				return
			}
		}

		index++
	}
	return nil
}
