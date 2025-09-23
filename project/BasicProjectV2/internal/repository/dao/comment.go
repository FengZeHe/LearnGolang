package dao

import (
	"database/sql"
	"errors"
	"log"

	"github.com/basicprojectv2/internal/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentDao interface {
	InsertComment(req domain.Comment) (err error)
	QueryComment(ctx *gin.Context, aid string) (comments []domain.Comment, err error)
	DeleteComment(ctx *gin.Context, id string) (err error)
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
		commentTemp := domain.Comment{
			Uid:     req.Uid,
			Aid:     req.Aid,
			Pid:     req.Pid,
			Rid:     sql.NullInt64{Valid: false},
			Content: req.Content,
		}
		if err = tx.Table("comment").Create(&commentTemp).Error; err != nil {
			log.Println("create comment error", err)
			return err
		}

		// 如果自身是顶级评论，那么rootid = 自身id
		if !req.Pid.Valid {
			commentTemp.Rid = sql.NullInt64{Int64: int64(commentTemp.Id), Valid: true}
		} else {
			// 把父级的root id拿过来
			var parent domain.Comment
			if err = tx.Table("comment").Where("id = ?", req.Pid.Int64).Find(&parent).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					log.Println("parent comment not found", err)
					return err
				}
				log.Println("find parent comment error", err)
				return err
			}
			commentTemp.Rid = parent.Rid
		}
		// 3. 确定rid，写回去
		return tx.Table("comment").Model(&domain.Comment{}).
			Where("id = ?", commentTemp.Id).Update("rid", commentTemp.Rid).Error
	})
}

func (c *GormCommentDAO) DeleteComment(ctx *gin.Context, id string) (err error) {
	return c.db.Transaction(func(tx *gorm.DB) error {
		var ids []int
		if err = tx.Table("comment").Model(&domain.Comment{}).Where("rid = ?", id).Pluck("id", &ids).Error; err != nil {
			log.Println("query comment error", err)
			return err
		}

		if len(ids) > 0 {
			if err = tx.Table("comment").Delete(&domain.Comment{}, ids).Error; err != nil {
				log.Println("delete comment error", err)
				return err
			}
		}
		//删除根评论
		if err = tx.Table("comment").Delete(&domain.Comment{}, id).Error; err != nil {
			log.Println("delete root comment error", err)
			return err
		}
		return nil
	})
}
