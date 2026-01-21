package domain

type Task struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `json:"name" gorm:"column:name"`
	CronExpr  string `json:"cronExpr" gorm:"column:cron_expr"`
	TaskName  string `json:"taskName" gorm:"column:task_name"`
	TaskType  string `json:"taskType" gorm:"column:task_type"`
	Status    int    `json:"status" gorm:"column:status"`
	CreatedAt string `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt string `json:"updatedAt" gorm:"column:updated_at"`
}

func (Task) TableName() string {
	return "tb_tasks"
}

type DeleteTaskReq struct {
	ID uint `json:"id"`
}

func (DeleteTaskReq) TableName() string {
	return "tb_tasks"
}

type TaskFilterReq struct {
	Status int `json:"status" default:"0"`
}

func (TaskFilterReq) TableName() string {
	return "tb_tasks"
}

type AddTaskReq struct {
	Name     string `json:"name"`
	CronExpr string `json:"cronExpr"`
	TaskName string `json:"taskName"`
	TaskType string `json:"taskType"`
}
