package domain

type Role struct {
	ID       string `json:"id"`
	RoleName string `json:"role_name" gorm:"column:role_name"`
	Desc     string `json:"desc" gorm:"column:desc"`
	Status   string `json:"status" gorm:"column:status"`
}
