package domain

type FollowReq struct {
	Uid    string `json:"uid"`
	Follow int    `json:"follow"`
}

type UserFollow struct {
	ID         string `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	FolloweeId string `gorm:"column:followee_id" json:"followee_id"`
	FollowerId string `gorm:"column:follower_id" json:"follower_id"`
	Status     string `gorm:"column:status" json:"status"`
	CreatedAt  string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  string `gorm:"column:updated_at" json:"updated_at"`
}

func (UserFollow) TableName() string {
	return "user_follow"
}
