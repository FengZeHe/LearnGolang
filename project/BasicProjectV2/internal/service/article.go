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
	GetArticlesByID(ctx context.Context, req domain.QueryArticlesByIDReq) (domain.Article, error)
	GetHotList(ctx context.Context) ([]domain.ArticleWithScores, error)
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

func (s *articleService) GetArticlesByID(ctx context.Context, req domain.QueryArticlesByIDReq) (l domain.Article, err error) {
	article, err := s.repo.GetArticleByID(ctx, req.ID)
	if err != nil {
		return domain.Article{}, err
	}
	return article, nil
}

func (s *articleService) GetHotList(ctx context.Context) ([]domain.ArticleWithScores, error) {
	list, err := s.repo.GetHotList(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}
