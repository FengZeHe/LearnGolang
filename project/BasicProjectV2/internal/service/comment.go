package service

import (
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository"
	"github.com/gin-gonic/gin"
)

type commentService struct {
	repo repository.CommentRepository
}

type CommentService interface {
	AddComment(ctx *gin.Context, req domain.AddCommentReq) (err error)
	GetComment(ctx *gin.Context, aid string) (comments []domain.Comment, err error)
	DeleteComment(ctx *gin.Context, aid string) (err error)
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentService{
		repo: repo,
	}
}

func (c *commentService) AddComment(ctx *gin.Context, req domain.AddCommentReq) (err error) {
	if err = c.repo.AddComment(ctx, req); err != nil {
		return err
	}
	return nil
}

func (c *commentService) GetComment(ctx *gin.Context, aid string) (comments []domain.Comment, err error) {
	comments, err = c.repo.GetComment(ctx, aid)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *commentService) DeleteComment(ctx *gin.Context, id string) (err error) {
	if err = c.repo.DeleteComment(ctx, id); err != nil {
		return err
	}
	return nil
}
