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
	GetArticles(ctx context.Context) ([]domain.ArticleResponse, error)
}

func NewArticleRepository(dao dao.ArticleDAO) ArticleRepository {
	return &articleRepository{articleDAO: dao}
}

func (a *articleRepository) GetArticles(ctx context.Context) (l []domain.ArticleResponse, err error) {
	res, err := a.articleDAO.GetArticles(ctx)
	if err != nil {
		return l, err
	}
	l = ArticleToEntity(res)
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
