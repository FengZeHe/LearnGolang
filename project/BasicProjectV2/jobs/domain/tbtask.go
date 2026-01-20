package domain

type Task struct {
	ID        int    `json:"id" Gorm:"primary_key;"`
	Name      string `json:"name" Gorm:"column:name"`
	CornExpr  string `json:"cornExpr" Gorm:"column:cron_expr"`
	TaskName  string `json:"taskName" Gorm:"column:task_name"`
	TaskType  string `json:"taskType" Gorm:"column:task_type"`
	Status    int    `json:"status" Gorm:"column:status"`
	CreatedAt string `json:"createdAt" Gorm:"column:created_at"`
	UpdatedAt string `json:"updatedAt" Gorm:"column:updated_at"`
}

func (Task) TableName() string {
	return "tb_tasks"
}

type DeleteTaskReq struct {
	ID int `json:"id"`
}

func (DeleteTaskReq) TableName() string {
	return "tb_tasks"
}

type TaskFilterReq struct {
	Status int `json:"status" default:"-1"`
}

func (TaskFilterReq) TableName() string {
	return "tb_tasks"
}
