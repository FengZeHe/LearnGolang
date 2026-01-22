package scheduler

import (
	"context"
	"log"
	"sync"

	"github.com/basicprojectv2/jobs"
	"github.com/basicprojectv2/jobs/domain"
	"github.com/basicprojectv2/jobs/repository/dao"
	"github.com/robfig/cron/v3"
)

type CronScheduler struct {
	cron      *cron.Cron
	taskDAO   dao.TaskDAO
	registry  *jobs.TaskRegistry
	running   bool
	mu        sync.RWMutex
	taskEntry map[uint]cron.EntryID // 任务ID -> Cron EntryID
}

// 创建Cron调度器
func NewCronScheduler(taskDAO dao.TaskDAO, registry *jobs.TaskRegistry) *CronScheduler {
	return &CronScheduler{
		cron:      cron.New(cron.WithSeconds()),
		taskDAO:   taskDAO,
		registry:  registry,
		running:   false,
		taskEntry: make(map[uint]cron.EntryID),
	}
}

// 启动调度器，加载正在执行的任务
func (s *CronScheduler) Start(ctx context.Context) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.running {
		return nil
	}

	acTasklist, acErr := s.taskDAO.FindActiveTasks()
	if acErr != nil {
		log.Println("Get Active Tasks Error:", acErr)
	}

	log.Println("Add Tasks:", len(acTasklist))
	for _, task := range acTasklist {
		if atErr := s.addTaskToCron(&task); atErr != nil {
			log.Println("Add Task Error:", err)
			return err
		}
		//log.Println("获取到", task.ID, task.Name, task.TaskType, task.CronExpr)
	}

	s.cron.Start()
	s.running = true

	go func() {
		<-ctx.Done()
		s.Stop()
	}()

	return nil

}

// Stop 停止调度器
func (s *CronScheduler) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.running {
		return
	}

	s.cron.Stop()
	s.running = false
	s.taskEntry = make(map[uint]cron.EntryID)
}

func (s *CronScheduler) RemoveTask(taskID uint) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 检查任务是否存在
	entryID, exists := s.taskEntry[taskID]
	if !exists {
		log.Println("任务不存在", taskID)
		return nil
	}
	s.cron.Remove(entryID)
	delete(s.taskEntry, taskID) // 删除map中的字段

	log.Println("已停止任务", taskID)
	return nil
}

// 添加Task到Cron中
func (s *CronScheduler) addTaskToCron(task *domain.Task) (err error) {

	// 创建cron解析器
	parser := cron.NewParser(
		cron.SecondOptional |
			cron.Minute |
			cron.Hour |
			cron.Dom |
			cron.Month |
			cron.Dow |
			cron.Descriptor,
	)

	schedule, err := parser.Parse(task.CronExpr)
	if err != nil {
		log.Println("cron expr error: %v\n", err)
		return err
	}

	// 获取正在执行的TaskName
	executor, exists := s.registry.Get(task.TaskName)
	if !exists {
		log.Println("该任务没有正在执行：", task.TaskName)
		return
	}

	entryID := s.cron.Schedule(schedule, cron.FuncJob(func() {
		s.execTask(task, executor)
	}))

	s.taskEntry[task.ID] = entryID
	return nil
}

// 执行任务
func (s *CronScheduler) execTask(task *domain.Task, exec jobs.TaskExecutor) {
	log.Println("开始执行任务:", task.Name)
	if err := s.taskDAO.UpdateTask(*task); err != nil {
		log.Println("Update Task Error:", err)
		return
	}

	if err := exec(task.ID); err != nil {
		log.Println("exec Task Error:", err)
		return
	}
}
