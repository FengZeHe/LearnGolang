package jobs

import "sync"

type TaskExecutor func(taskID uint) (err error)

type TaskRegistry struct {
	executors map[string]TaskExecutor
	mu        sync.RWMutex
}

func NewTaskRegistry() *TaskRegistry {
	return &TaskRegistry{
		executors: make(map[string]TaskExecutor),
	}
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
