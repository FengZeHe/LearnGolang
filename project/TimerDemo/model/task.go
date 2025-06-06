package model

import "time"

// TbTask Model
type TbTasks struct {
	ID        uint      `json:"id" `
	Name      string    `json:"name" gorm:"column:name"`
	CronExpr  string    `json:"cronExpr" gorm:"column:cron_expr"`
	TaskType  string    `json:"taskType" gorm:"column:task_type"` // 任务类型，脚本/函数 script/fun
	TaskName  string    `json:"taskName" gorm:"column:task_name"` // 执行任务名称
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

type TbTaskLog struct {
	ID       int    `json:"id"`
	TaskID   int    `json:"taskId"`
	ExecTime string `json:"execTime"`
}

type AddTaskReq struct {
	Name     string `json:"name"`
	CronExpr string `json:"cronExpr"`
	TaskType string `json:"taskType"`
	TaskName string `json:"taskName"`
}

type UpdateTaskReq struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}
