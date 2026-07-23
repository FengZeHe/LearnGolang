package dao

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/basicprojectv2/internal/domain"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type GORMArticle struct {
	db  *gorm.DB
	rdb redis.Cmdable
	esc *elasticsearch.Client
}

type ArticleDAO interface {
	InsertArticle(ctx context.Context, a domain.Article) error
	UpdateArticleByID(ctx context.Context, a domain.Article) error
	GetArticles(ctx context.Context, pageIndex, pageSize int) (domain.ArticlesDAOResponse, error)
	GetArticleByID(ctx context.Context, id string) (domain.Article, error)
	GetArticlesByUserID(ctx context.Context, pageIndex, pageSize int, userID string) (domain.ArticlesDAOResponse, error)
	AddArticleCount(ctx context.Context, id string) error
	GetHosList(ctx context.Context, key string) ([]domain.ArticleWithScores, error)
	SearchByKeyword(ctx context.Context, keyword string, pageIndex, pageSize int) (map[string]interface{}, error)
	SyncArticles() error
}

func NewArticleDAO(db *gorm.DB, rdb redis.Cmdable, esc *elasticsearch.Client) ArticleDAO {
	return &GORMArticle{
		db:  db,
		rdb: rdb,
		esc: esc,
	}
}

func (dao *GORMArticle) SearchByKeyword(ctx context.Context, keyword string, pageIndex, pageSize int) (map[string]interface{}, error) {
	data := make(map[string]interface{})

	// 构建查询DSL

	from := (pageIndex - 1) * pageSize
	query := map[string]interface{}{
		"from": from,
		"size": pageSize,
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  keyword,
				"fields": []string{"title^3", "content^1", "tags"},
				"type":   "best_fields",
			},
		},
		"highlight": map[string]interface{}{
			"fields": map[string]interface{}{
				"title":   map[string]interface{}{"number_of_fragments": 0},
				"content": map[string]interface{}{"fragment_size": 10, "number_of_fragments": 3},
			},
		},
	}

	// 执行搜索
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Println("New Encoder Error", err)
		return data, err
	}

	res, err := dao.esc.Search(
		dao.esc.Search.WithIndex("articles"),
		dao.esc.Search.WithBody(&buf),
	)
	if err != nil {
		log.Println("es search error", err)
		return data, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Println("es search Decode error", err)
		return data, err
	}
	log.Println("result ==>", result)

	// 提取 hits
	hits, ok := result["hits"].(map[string]interface{})
	if !ok {
		return data, errors.New("SearchError")
	}

	// 提取总条数
	total := int64(0)
	if totalObj, ok := hits["total"].(map[string]interface{}); ok {
		if val, ok := totalObj["value"].(float64); ok {
			total = int64(val)
		}
	}

	// 提取文档列表 + 高亮
	var list []map[string]interface{}
	if hitsList, ok := hits["hits"].([]interface{}); ok {
		for _, hit := range hitsList {
			hitMap := hit.(map[string]interface{})
			source := hitMap["_source"].(map[string]interface{})

			// 合并高亮字段
			if highlight, ok := hitMap["highlight"]; ok {
				for k, v := range highlight.(map[string]interface{}) {
					source["highlight_"+k] = v
				}
			}
			list = append(list, source)
		}
	}

	data["total"] = total
	data["list"] = list
	data["pageSize"] = pageSize
	data["pageIndex"] = pageIndex

	return data, nil
}

func (dao *GORMArticle) SyncArticles() error {
	const indexName = "articles"
	const batchSize = 100
	var lastID string
	var totalSynced int64

	for {
		var articles []domain.Article
		query := dao.db.Table("article").Order("id ASC").Limit(batchSize)
		if lastID != "" {
			query = query.Where("id > ?", lastID)
		}
		if err := query.Find(&articles).Error; err != nil {
			return fmt.Errorf("query article ERROR: %w", err)
		}

		if len(articles) == 0 {
			break
		}

		// 构建 bulk 请求
		var buf bytes.Buffer
		for _, art := range articles {
			esDoc := domain.ArticleWithES{
				ID:        art.ID,
				Title:     art.Title,
				Content:   art.Content,
				Tags:      art.Tags,
				Status:    art.Status,
				CreatedAt: art.CreatedAt,
				UpdatedAt: art.UpdatedAt,
			}

			// 添加 index action 元数据
			meta := map[string]interface{}{
				"index": map[string]interface{}{
					"_index": indexName,
					"_id":    art.ID,
				},
			}
			if err := json.NewEncoder(&buf).Encode(meta); err != nil {
				log.Printf("序列化 meta 失败 ID=%s: %v", art.ID, err)
				continue
			}

			// 添加文档数据
			if err := json.NewEncoder(&buf).Encode(esDoc); err != nil {
				log.Printf("序列化文章失败 ID=%s: %v", art.ID, err)
				continue
			}

			lastID = art.ID
		}

		// 执行 bulk 请求
		res, err := dao.esc.Bulk(bytes.NewReader(buf.Bytes()), dao.esc.Bulk.WithIndex(indexName))
		if err != nil {
			log.Printf("执行 bulk 请求失败: %v", err)
			return err
		}

		if res.IsError() {
			log.Printf("Bulk 请求错误: %s", res.Status())
		}

		// 解析响应检查错误
		var bulkRes map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&bulkRes); err != nil {
			log.Printf("解析 bulk 响应失败: %v", err)
		} else if hasErrors, ok := bulkRes["errors"].(bool); ok && hasErrors {
			log.Printf("Bulk 部分操作失败: %v", bulkRes)
		}

		_ = res.Body.Close()

		totalSynced += int64(len(articles))
		log.Printf("已同步 %d 篇文章", totalSynced)

		if len(articles) < batchSize {
			break
		}
	}

	log.Printf("同步完成，共 %d 篇文章", totalSynced)
	return nil
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
