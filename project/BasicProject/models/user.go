package models

type RegisterForm struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type LoginForm struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Id       string `json:"id" `
	Email    string `json:"email" gorm:"type:varchar(255);unique"`
	Password string `json:"-" gorm:"type:varchar(255)"`
	Phone    string `json:"phone" gorm:"type:varchar(255);default:null"`
	Birthday int64  `json:"birthday"  gorm:"column:birthday;default:null"`
	Nickname string `json:"nickname" gorm:"type:varchar(255);column:nickname;default:null"`
	Aboutme  string `json:"aboutme" gorm:"type:varchar(255);column:about_me;default:null"`
	Ctime    int64  `json:"ctime"`
	Utime    int64  `json:"utime" `
}
