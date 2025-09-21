package domain

type Comment struct {
	Id        int64  `gorm:"column:id" json:"id"`
	Uid       int64  `gorm:"column:uid" json:"uid"` // 发表评论用户的ID
	Pid       int64  `gorm:"column:pid" json:"pid"` // 评论的父ID
	Rid       int64  `gorm:"column:rid" json:"rid"` // 评论的ROOT ID
	Content   string `gorm:"column:content"  json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
