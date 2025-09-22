package service

import "github.com/basicprojectv2/internal/repository"

type commentService struct {
	repo repository.CommentRepository
}

type CommentService interface {
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentService{
		repo: repo,
	}
}
