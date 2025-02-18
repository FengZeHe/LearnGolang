package dao

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"gorm.io/gorm"
	"math"
	"time"
)

type GORMArticle struct {
	db *gorm.DB
}

type ArticleDAO interface {
	InsertArticle(ctx context.Context, a domain.Article) error
	UpdateArticleByID(ctx context.Context, a domain.Article) error
	GetArticles(ctx context.Context, pageIndex, pageSize int) (domain.ArticlesDAOResponse, error)
	GetArticlesByID(ctx context.Context, pageIndex, pageSize int, userID string) (domain.ArticlesDAOResponse, error)
	AddArticleCount(ctx context.Context, id string) error
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

func (dao *GORMArticle) GetArticles(ctx context.Context, pageIndex, pageSize int) (a domain.ArticlesDAOResponse, err error) {
	var data []domain.Article

	// 计算偏移量
	offset := (pageIndex - 1) * pageSize

	// 查询总记录数
	var totalCount int64
	dao.db.Model(&domain.Article{}).Table("article").Count(&totalCount)

	//执行分页查询
	if err = dao.db.WithContext(ctx).Table("article").Limit(pageSize).Offset(offset).Order("created_at desc").Find(&data).Error; err != nil {
		return a, err
	}

	// 计算总页数
	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	if totalCount == 0 || totalPages == 0 {
		totalPages = 1
	}

	a.Articles = data
	a.PageIndex = pageIndex
	a.PageCount = totalPages
	a.TotalCount = totalCount
	return a, nil
}

func (dao *GORMArticle) GetArticlesByID(ctx context.Context, pageIndex, pageSize int, userID string) (a domain.ArticlesDAOResponse, err error) {
	var data []domain.Article

	// 计算偏移量
	offset := (pageIndex - 1) * pageSize

	// 查询总记录数
	var totalCount int64
	dao.db.Model(&domain.Article{}).Table("article").Where("author_id = ?", userID).Count(&totalCount)

	//执行分页查询
	if err = dao.db.WithContext(ctx).Table("article").Where("author_id = ?", userID).Limit(pageSize).Offset(offset).Order("created_at desc").Find(&data).Error; err != nil {
		return a, err
	}

	// 计算总页数
	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	if totalCount == 0 || totalPages == 0 {
		totalPages = 1
	}

	a.Articles = data
	a.PageIndex = pageIndex
	a.PageCount = totalPages
	a.TotalCount = totalCount
	return a, nil
}

func (dao *GORMArticle) AddArticleCount(ctx context.Context, id string) (err error) {
	if err = dao.db.Model(&domain.Article{}).Table("article").Where("id = ?", id).Update("`read`", gorm.Expr("`read` + 1")).Error; err != nil {
		return err
	}
	return nil
}
