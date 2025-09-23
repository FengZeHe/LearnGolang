package domain

import "database/sql"

type Comment struct {
	Id        int           `gorm:"column:id;primaryKey" json:"id"`
	Uid       int64         `gorm:"column:uid" json:"uid"` // 发表评论用户的ID
	Pid       sql.NullInt64 `gorm:"column:pid" json:"pid"` // 评论的父ID
	Rid       sql.NullInt64 `gorm:"column:rid" json:"rid"` // 评论的ROOT ID
	Aid       int64         `gorm:"column:aid" json:"aid"` // 评论的文章ID
	Content   string        `gorm:"column:content"  json:"content"`
	CreatedAt string        `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt string        `json:"updatedAt"`
}

type AddCommentReq struct {
	Pid     string `json:"pid"`
	Aid     string `json:"aid"`
	Content string `json:"content"`
}

type CommentResp struct {
	Id        int           `gorm:"column:id;primaryKey" json:"id"`
	Uid       int64         `gorm:"column:uid" json:"uid"` // 发表评论用户的ID
	Pid       sql.NullInt64 `gorm:"column:pid" json:"pid"` // 评论的父ID
	Rid       sql.NullInt64 `gorm:"column:rid" json:"rid"` // 评论的ROOT ID
	Aid       int64         `gorm:"column:aid" json:"aid"` // 评论的文章ID
	Content   string        `gorm:"column:content"  json:"content"`
	CreatedAt string        `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt string        `json:"updatedAt"`
	Email     string        ` json:"email"`
}
