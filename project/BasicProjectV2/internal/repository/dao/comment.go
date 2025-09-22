package dao

import (
	"errors"

	"github.com/basicprojectv2/internal/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentDao interface {
	InsertComment(req domain.Comment) (err error)
	QueryComment(ctx *gin.Context, aid string) (comments []domain.Comment, err error)
	DeleteComment(ctx *gin.Context, aid string) (err error)
}

type GormCommentDAO struct {
	db *gorm.DB
}

func NewCommentDao(db *gorm.DB) CommentDao {
	return &GormCommentDAO{
		db: db,
	}
}

func (c *GormCommentDAO) QueryComment(ctx *gin.Context, aid string) (comments []domain.Comment, err error) {
	if err = c.db.Table("comment").Where("aid = ?", aid).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
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

func (c *GormCommentDAO) DeleteComment(ctx *gin.Context, aid string) (err error) {
	/*
		todo 除了删除自身评论，还要删除底下的子评论(pid = 该id)
	*/
	return nil
}
