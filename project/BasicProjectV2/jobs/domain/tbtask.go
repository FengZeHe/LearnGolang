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

type HostScoreCalc struct {
	ReadCount    int    `gorm:"column:read_count"`
	LikeCount    int    `gorm:"column:like_count"`
	CollectCount int    `gorm:"column:collect_count"`
	CreatedAt    string `gorm:"column:created_at"`
}

func (HostScoreCalc) TableName() string {
	return "interactive"
}

type Article struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Status     string `json:"status"`
	Read       int    `json:"read"`
	AuthorName string `json:"authorName"`
	AuthorID   string `json:"authorId"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"-"`
}

func (Article) TableName() string {
	return "article"
}

type Interactive struct {
	ID           string `json:"id" gorm:"primary_key;autoIncrement;"`
	Aid          string `json:"aid"`
	ReadCount    int    `json:"readCount" gorm:"column:read_count"`
	LikeCount    int    `json:"likeCount" gorm:"column:like_count"`
	CollectCount int    `json:"collectCount" gorm:"column:collect_count"`
	CTime        string `json:"-" gorm:"column:ctime"`
	UTime        string `json:"-" gorm:"column:utime"`
}

func (Interactive) TableName() string {
	return "interactive"
}

type ArticleWithInteractive struct {
	ID           string `json:"id" gorm:"primary_key;autoIncrement;"`
	Title        string `gorm:"column:title"`
	ReadCount    int    `json:"readCount" gorm:"column:read_count"`
	LikeCount    int    `json:"likeCount" gorm:"column:like_count"`
	CollectCount int    `json:"collectCount" gorm:"column:collect_count"`
	CreatedAt    string `json:"createdAt" gorm:"column:created_at"`
}
