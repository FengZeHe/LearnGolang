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
