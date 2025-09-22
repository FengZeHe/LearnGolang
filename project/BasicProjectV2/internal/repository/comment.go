package repository

import (
	"strconv"
	"time"

	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository/dao"
	"github.com/gin-gonic/gin"
)

type commentRepository struct {
	dao dao.CommentDao
}

type CommentRepository interface {
	AddComment(ctx *gin.Context, req domain.AddCommentReq) error
}

func NewCommentRepository(dao dao.CommentDao) CommentRepository {
	return &commentRepository{
		dao: dao,
	}
}

func (c *commentRepository) AddComment(ctx *gin.Context, req domain.AddCommentReq) (err error) {
	userid, _ := ctx.Get("userid")
	strUserid := userid.(string)
	comment := domain.Comment{
		Uid:       StringToInt64(strUserid),
		Aid:       StringToInt64(req.Aid),
		Content:   req.Content,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err = c.dao.InsertComment(comment); err != nil {
		return err
	}
	return nil
}

func StringToInt64(target string) (output int64) {
	output, _ = strconv.ParseInt(target, 10, 64)
	return output
}
