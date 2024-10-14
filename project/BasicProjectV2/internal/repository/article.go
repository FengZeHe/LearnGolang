package repository

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository/dao"
)

type articleRepository struct {
	articleDAO dao.ArticleDAO
}

type ArticleRepository interface {
	GetArticles(ctx context.Context, req domain.QueryArticlesReq) (domain.ArticleRepoResponse, error)
	GetAuthorArticles(ctx context.Context, req domain.QueryAuthorArticlesReq, userid string) (domain.ArticleRepoResponse, error)
}

func NewArticleRepository(dao dao.ArticleDAO) ArticleRepository {
	return &articleRepository{articleDAO: dao}
}

func (a *articleRepository) GetArticles(ctx context.Context, req domain.QueryArticlesReq) (l domain.ArticleRepoResponse, err error) {
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.PageIndex <= 0 {
		req.PageIndex = 1
	}

	res, err := a.articleDAO.GetArticles(ctx, req.PageIndex, req.PageSize)
	if err != nil {
		return l, err
	}
	t := ArticleToEntity(res.Articles)
	l.Articles = t
	l.PageIndex = res.PageIndex
	l.PageCount = res.PageCount
	l.TotalCount = int(res.TotalCount)

	return l, err
}

func (a *articleRepository) GetAuthorArticles(ctx context.Context, req domain.QueryAuthorArticlesReq, userid string) (l domain.ArticleRepoResponse, err error) {
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.PageIndex <= 0 {
		req.PageIndex = 1
	}

	res, err := a.articleDAO.GetArticlesByID(ctx, req.PageIndex, req.PageSize, userid)
	if err != nil {
		return l, err
	}
	t := ArticleToEntity(res.Articles)
	l.Articles = t
	l.PageIndex = res.PageIndex
	l.PageCount = res.PageCount
	l.TotalCount = int(res.TotalCount)

	return l, err
}

func ArticleToEntity(origin []domain.Article) (target []domain.ArticleResponse) {
	for _, v := range origin {
		target = append(target, domain.ArticleResponse{
			ID:         v.ID,
			Title:      v.Title,
			Content:    v.Content,
			AuthorName: v.AuthorName,
			CreatedAt:  v.CreatedAt,
		})
	}
	return target
}
