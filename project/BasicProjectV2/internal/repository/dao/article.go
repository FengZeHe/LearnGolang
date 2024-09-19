package dao

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"gorm.io/gorm"
	"time"
)

type GORMArticle struct {
	db *gorm.DB
}

type ArticleDAO interface {
	InsertArticle(ctx context.Context, a domain.Article) error
	UpdateArticleByID(ctx context.Context, a domain.Article) error
	GetArticles(ctx context.Context) ([]domain.Article, error)
}

func NewArticleDAO(db *gorm.DB) ArticleDAO {
	return &GORMArticle{
		db: db,
	}
}

func (dao *GORMArticle) InsertArticle(ctx context.Context, a domain.Article) (err error) {
	a.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	if err = dao.db.WithContext(ctx).Table("article").Create(&a).Error; err != nil {
		return err
	}
	return nil
}

func (dao *GORMArticle) UpdateArticleByID(ctx context.Context, a domain.Article) (err error) {
	if err = dao.db.WithContext(ctx).Table("article").Save(&a).Error; err != nil {
		return err
	}
	return nil
}

func (dao *GORMArticle) GetArticles(ctx context.Context) (a []domain.Article, err error) {
	if err = dao.db.WithContext(ctx).Table("article").Order("created_at desc").Find(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}
