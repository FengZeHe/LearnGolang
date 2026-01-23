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
	GetArticleByID(ctx context.Context, articleID string) (domain.Article, error)
	GetAuthorArticles(ctx context.Context, req domain.QueryAuthorArticlesReq, userid string) (domain.ArticleRepoResponse, error)
	AddArticleReadCount(ctx context.Context, id string) error
	GetHotList(ctx context.Context) ([]domain.ArticleWithScores, error)
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

func (a *articleRepository) GetArticleByID(ctx context.Context, articleID string) (domain.Article, error) {
	article, err := a.articleDAO.GetArticleByID(ctx, articleID)
	if err != nil {
		return domain.Article{}, err
	}

	return article, nil
}

func (a *articleRepository) GetAuthorArticles(ctx context.Context, req domain.QueryAuthorArticlesReq, userid string) (l domain.ArticleRepoResponse, err error) {
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.PageIndex <= 0 {
		req.PageIndex = 1
	}

	res, err := a.articleDAO.GetArticlesByUserID(ctx, req.PageIndex, req.PageSize, userid)
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

func (a *articleRepository) AddArticleReadCount(ctx context.Context, id string) (err error) {
	if err = a.articleDAO.AddArticleCount(ctx, id); err != nil {
		return err
	}
	return nil
}

func ArticleToEntity(origin []domain.Article) (target []domain.ArticleResponse) {
	for _, v := range origin {
		target = append(target, domain.ArticleResponse{
			ID:         v.ID,
			Title:      v.Title,
			Content:    v.Content,
			AuthorName: v.AuthorName,
			CreatedAt:  v.CreatedAt,
			Read:       v.Read,
		})
	}
	return target
}

const HostListScoreKey = "hotlist/articles/score/"

func (a *articleRepository) GetHotList(ctx context.Context) (res []domain.ArticleWithScores, err error) {
	res, err = a.articleDAO.GetHosList(ctx, HostListScoreKey)
	if err != nil {
		return res, err
	}
	return res, nil
}
