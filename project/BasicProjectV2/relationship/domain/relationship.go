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

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Phone    string `json:"phone"`
	Birthday int    `json:"birthday"`
	Nickname string `json:"nickname"`
	Aboutme  string `json:"aboutme"`
	Role     string `json:"role"`
}

func (User) TableName() string {
	return "users"
}

type Relationship struct {
	ID          string `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Uid         string `gorm:"column:uid;" json:"uid"`
	FolloweeNum int64  `gorm:"column:followee_num;CHECK(followee_num >= 0)" json:"followee_num"`
	FollowerNum int64  `gorm:"column:follower_num;CHECK(follower_num >= 0)" json:"follower_num"`
	CreatedAt   string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   string `gorm:"column:updated_at" json:"updated_at"`
}

func (Relationship) TableName() string {
	return "relationship_record"
}
