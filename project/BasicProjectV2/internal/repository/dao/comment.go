package dao

import "gorm.io/gorm"

type CommentDao interface {
}

type GormCommentDAO struct {
	db *gorm.DB
}

func NewCommentDao(db *gorm.DB) CommentDao {
	return &GormCommentDAO{
		db: db,
	}
}
