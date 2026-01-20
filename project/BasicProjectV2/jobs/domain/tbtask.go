package domain

type TbTask struct {
	ID        int    `json:"id" Gorm:"primary_key;"`
	Name      string `json:"name" Gorm:"column:name"`
	CornExpr  string `json:"cornExpr" Gorm:"column:cron_expr"`
	TaskName  string `json:"taskName" Gorm:"column:task_name"`
	TaskType  string `json:"taskType" Gorm:"column:task_type"`
	Status    string `json:"status" Gorm:"column:status"`
	CreatedAt string `json:"createdAt" Gorm:"column:created_at"`
	UpdatedAt string `json:"updatedAt" Gorm:"column:updated_at"`
}

func (TbTask) TableName() string {
	return "tb_tasks"
}

type AddTask struct {
	Name     string `json:"name" Gorm:"column:name"`
	CornExpr string `json:"cornExpr" Gorm:"column:cron_expr"`
	TaskName string `json:"taskName" Gorm:"column:task_name"`
	TaskType string `json:"taskType" Gorm:"column:task_type"`
}

func (AddTask) TableName() string {
	return "tb_tasks"
}
