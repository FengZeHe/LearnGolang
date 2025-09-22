package dao

import (
	"errors"

	"github.com/basicprojectv2/internal/domain"
	"gorm.io/gorm"
)

type CommentDao interface {
	InsertComment(req domain.Comment) (err error)
}

type GormCommentDAO struct {
	db *gorm.DB
}

func NewCommentDao(db *gorm.DB) CommentDao {
	return &GormCommentDAO{
		db: db,
	}
}

func (c *GormCommentDAO) InsertComment(req domain.Comment) (err error) {
	article := domain.Article{}
	var count int64
	if err = c.db.Table("article").Find(&article).Where("id = ?", req.Aid).Count(&count).Error; err != nil {
		return err
	}

	if count <= 0 {
		return errors.New("NoArticle")
	} else {
		if err = c.db.Table("comment").Create(&req).Error; err != nil {
			return err
		}
	}

	return nil
}
