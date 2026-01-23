package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/service"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

type ArticleHandler struct {
	svc service.ArticleService
	//readProducer article.Producer
}

func NewArticleHandler(svc service.ArticleService) *ArticleHandler {
	return &ArticleHandler{svc: svc}
}

func (r *ArticleHandler) RegisterRoutes(server *gin.Engine, loginCheck gin.HandlerFunc) {
	rg := server.Group("/v2/article/")
	rg.Use(loginCheck)

	rg.POST("/getArticles", r.GetArticles)
	rg.POST("/getAuthorArticles", r.GetAuthorArticles)
	rg.POST("/getArticleByID", r.GetArticleByID)
	//rg.POST("/addReadCount", r.AddReadCount)
	rg.GET("/hotList", r.GetHotList)
}

func (r *ArticleHandler) GetArticles(c *gin.Context) {
	_, span := otel.Tracer("gin-service").Start(c.Request.Context(), "handleGetArticle")
	defer span.End()

	req := domain.QueryArticlesReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		span.RecordError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求参数错误",
		})
		return
	}
	_, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户未登录",
		})
		return
	}
	data, err := r.svc.GetArticles(c, req)
	if err != nil {
		log.Println(err)
		span.RecordError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	}
	span.SetStatus(codes.Ok, "Success")
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (r *ArticleHandler) GetArticleByID(c *gin.Context) {
	req := domain.QueryArticlesByIDReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求参数错误",
		})
		return
	}
	_, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户未登录",
		})
		return
	}

	data, err := r.svc.GetArticlesByID(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "返回错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (r *ArticleHandler) GetAuthorArticles(c *gin.Context) {
	req := domain.QueryAuthorArticlesReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "请求参数错误",
		})
		return
	}
	userid, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户未登录",
		})
		return
	}
	userIDStr := fmt.Sprintf("%v", userid)
	data, err := r.svc.GetAuthorArticles(c, req, userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (r *ArticleHandler) AddReadCount(c *gin.Context) {
	req := domain.AddArticleCount{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	log.Println("Read Count + 1", req.ID)
	articleID, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	log.Println(articleID)

	// 发送阅读事件到kafka
	//evt := article.ReadEvent{
	//	Aid: articleID,
	//	Uid: c.GetInt64("userid"),
	//}
	//
	//if err := r.readProducer.ProduceReadEvent(evt); err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"msg": err.Error(),
	//	})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})

}

func (r *ArticleHandler) GetHotList(c *gin.Context) {
	list, err := r.svc.GetHotList(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": list,
	})
}
