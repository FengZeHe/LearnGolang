package dao

import (
	"context"
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/basicprojectv2/internal/domain"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type GORMArticle struct {
	db  *gorm.DB
	rdb redis.Cmdable
}

type ArticleDAO interface {
	InsertArticle(ctx context.Context, a domain.Article) error
	UpdateArticleByID(ctx context.Context, a domain.Article) error
	GetArticles(ctx context.Context, pageIndex, pageSize int) (domain.ArticlesDAOResponse, error)
	GetArticleByID(ctx context.Context, id string) (domain.Article, error)
	GetArticlesByUserID(ctx context.Context, pageIndex, pageSize int, userID string) (domain.ArticlesDAOResponse, error)
	AddArticleCount(ctx context.Context, id string) error
	GetHosList(ctx context.Context, key string) ([]domain.ArticleWithScores, error)
}

func NewArticleDAO(db *gorm.DB, rdb redis.Cmdable) ArticleDAO {
	return &GORMArticle{
		db:  db,
		rdb: rdb,
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

func (dao *GORMArticle) GetArticlesByUserID(ctx context.Context, pageIndex, pageSize int, userID string) (a domain.ArticlesDAOResponse, err error) {
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

func (dao *GORMArticle) GetArticleByID(ctx context.Context, id string) (a domain.Article, err error) {
	err = dao.db.WithContext(ctx).Table("article").Where("id = ?", id).First(&a).Error
	if err != nil {
		return a, err
	}
	return a, nil
}

func (dao *GORMArticle) AddArticleCount(ctx context.Context, id string) (err error) {
	if err = dao.db.Model(&domain.Article{}).Table("article").Where("id = ?", id).Update("`read`", gorm.Expr("`read` + 1")).Error; err != nil {
		return err
	}
	return nil
}

func (dao *GORMArticle) GetHosList(ctx context.Context, key string) (hostList []domain.ArticleWithScores, err error) {
	pipe := dao.rdb.Pipeline()
	zCmd := pipe.ZRevRangeWithScores(ctx, key, 0, 10)

	_, _ = pipe.Exec(ctx)

	zs, err := zCmd.Result()
	if err != nil {
		return nil, err
	}
	if len(zs) == 0 {
		return nil, nil
	}

	pipe2 := dao.rdb.Pipeline()
	titleCmds := make([]*redis.StringCmd, len(zs))
	for i, z := range zs {
		idStr := z.Member.(string)
		titleCmds[i] = pipe2.HGet(ctx, "hotlist/articles/"+idStr, "title")
	}
	_, err = pipe2.Exec(ctx)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	output := make([]domain.ArticleWithScores, 0, len(zs))
	for i, z := range zs {
		id, _ := strconv.ParseUint(z.Member.(string), 10, 64)
		output = append(output, domain.ArticleWithScores{
			ID:    strconv.FormatUint(id, 10),
			Title: titleCmds[i].Val(),
			Score: z.Score,
		})
	}
	hostList = output
	return hostList, nil
}
