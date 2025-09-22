package dao

import (
	"database/sql"
	"log"

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

	return c.db.Transaction(func(tx *gorm.DB) error {
		// todo 1.判断该评论是否为顶级评论(前面req可以判断)

		commentTemp := domain.Comment{
			Uid: req.Uid,
			Aid: req.Aid,
			Pid: req.Pid, //判断pid是否
			Rid: sql.NullInt64{Valid: false},
		}
		if err = tx.Table("comment").Create(&commentTemp).Error; err != nil {
			log.Println("create comment error", err)
			return err
		}

		// 如果存在pid
		if !req.Pid.Valid {
			commentTemp.Rid = sql.NullInt64{Int64: int64(commentTemp.Id), Valid: true}
		} else {
			// 把父级的root id拿过来
			var parent domain.Comment
			// todo 还要修改，当查不到record时 rid是0
			if err = tx.Table("comment").Where("id = ?", req.Pid).Find(&parent).Error; err != nil {
				log.Println("find parent comment error", err)
				return err
			}
			commentTemp.Rid = parent.Rid
		}

		// 3. 确定rid，写回去
		return tx.Model(&domain.Comment{}).
			Where("id = ?", commentTemp.Id).Update("rid", commentTemp.Rid).Error
	})
}

func (c *GormCommentDAO) DeleteComment(ctx *gin.Context, aid string) (err error) {
	/*
		todo 除了删除自身评论，还要删除底下的子评论(pid = 该id)
	*/
	return nil
}
