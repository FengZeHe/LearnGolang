package model

// TbTask Model
type TbTasks struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CronExpr  string `json:"cronExpr"`
	Status    int    `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type TbTaskLog struct {
	ID       int    `json:"id"`
	TaskID   int    `json:"taskId"`
	ExecTime string `json:"execTime"`
}

type AddTaskReq struct {
	Name     string `json:"name"`
	CronExpr string `json:"cronExpr"`
}
