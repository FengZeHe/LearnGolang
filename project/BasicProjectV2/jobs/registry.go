package jobs

import (
	"sync"

	"github.com/basicprojectv2/jobs/fun"
	"github.com/basicprojectv2/jobs/repository/dao"
	"github.com/redis/go-redis/v9"
)

type TaskExecutor func(taskID uint) (err error)

type TaskRegistry struct {
	executors map[string]TaskExecutor
	mu        sync.RWMutex
	taskDAO   dao.TaskDAO
	rdb       redis.Cmdable
}

func NewTaskRegistry(taskDAO dao.TaskDAO, rdb redis.Cmdable) *TaskRegistry {
	r := &TaskRegistry{
		executors: make(map[string]TaskExecutor),
		taskDAO:   taskDAO,
		rdb:       rdb,
	}
	// 注册任务
	r.Register("SayHi", fun.ExecTimeKeeping)

	// todo 闭包知识点
	r.Register("CalcHotList", func(taskID uint) error {
		return fun.CalcHotList(taskID, r.taskDAO, r.rdb)
	})
	return r
}

func (tr *TaskRegistry) Register(taskName string, executor TaskExecutor) {
	tr.mu.Lock()
	defer tr.mu.Unlock()
	tr.executors[taskName] = executor
}

// Get 获取任务执行函数
func (tr *TaskRegistry) Get(taskName string) (TaskExecutor, bool) {
	tr.mu.RLock()
	defer tr.mu.RUnlock()
	executor, exists := tr.executors[taskName]
	return executor, exists
}
