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
	GetArticles(ctx context.Context, req domain.QueryArticlesReq) (domain.ArticleRepoResponse, error)
	GetAuthorArticles(ctx context.Context, req domain.QueryAuthorArticlesReq, userid string) (domain.ArticleRepoResponse, error)
	AddArticleReadCount(ctx context.Context, id string) error
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

func (s *articleService) GetArticles(ctx context.Context, req domain.QueryArticlesReq) (l domain.ArticleRepoResponse, err error) {
	l, err = s.repo.GetArticles(ctx, req)
	if err != nil {
		return l, err
	}

	return l, err
}

func (s *articleService) GetAuthorArticles(ctx context.Context, req domain.QueryAuthorArticlesReq, userid string) (l domain.ArticleRepoResponse, err error) {
	l, err = s.repo.GetAuthorArticles(ctx, req, userid)
	if err != nil {
		return l, err
	}
	return l, nil
}

func (s *articleService) AddArticleReadCount(ctx context.Context, id string) (err error) {
	if err = s.repo.AddArticleReadCount(ctx, id); err != nil {
		return err
	}
	return nil
}
