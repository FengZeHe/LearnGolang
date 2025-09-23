package repository

import (
	"database/sql"
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
	GetComment(ctx *gin.Context, aid string) (comments []domain.Comment, err error)
	DeleteComment(ctx *gin.Context, id string) (err error)
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
		Pid:       StringToSqlNullInt64(req.Pid),
		Content:   req.Content,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err = c.dao.InsertComment(comment); err != nil {
		return err
	}
	return nil
}

func (c *commentRepository) GetComment(ctx *gin.Context, aid string) (comments []domain.Comment, err error) {
	comments, err = c.dao.QueryComment(ctx, aid)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *commentRepository) DeleteComment(ctx *gin.Context, id string) (err error) {
	return c.dao.DeleteComment(ctx, id)
}

func StringToSqlNullInt64(target string) (output sql.NullInt64) {
	// 如果pid没有填
	if target == "" {
		output.Valid = false
	} else {
		output.Valid = true
	}

	outputValue, _ := strconv.ParseInt(target, 10, 64)
	output.Int64 = outputValue
	return output
}
func StringToInt64(value string) (output int64) {
	output, _ = strconv.ParseInt(value, 10, 64)
	return output
}
