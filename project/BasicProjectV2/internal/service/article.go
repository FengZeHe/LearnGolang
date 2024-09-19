package service

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository"
)

type articleService struct {
	repo repository.ArticleRepository
}

type ArticleService interface {
	GetArticles(ctx context.Context) ([]domain.ArticleResponse, error)
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

func (s *articleService) GetArticles(ctx context.Context) (l []domain.ArticleResponse, err error) {
	l, err = s.repo.GetArticles(ctx)
	if err != nil {
		return nil, err
	}

	return l, err
}
