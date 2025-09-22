package domain

type Comment struct {
	Id        int64  `gorm:"column:id" json:"id"`
	Uid       int64  `gorm:"column:uid" json:"uid"` // 发表评论用户的ID
	Pid       int64  `gorm:"column:pid" json:"pid"` // 评论的父ID
	Rid       int64  `gorm:"column:rid" json:"rid"` // 评论的ROOT ID
	Aid       int64  `gorm:"column:aid" json:"aid"` // 评论的文章ID
	Content   string `gorm:"column:content"  json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type AddCommentReq struct {
	Uid     int64  `json:"uid"`
	Pid     int64  `json:"pid"`
	Aid     int64  `json:"aid"`
	Content string `json:"content"`
}
