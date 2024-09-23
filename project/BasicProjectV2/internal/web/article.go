package web

import (
	"fmt"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ArticleHandler struct {
	svc service.ArticleService
}

func NewArticleHandler(svc service.ArticleService) *ArticleHandler {
	return &ArticleHandler{svc: svc}
}

func (r *ArticleHandler) RegisterRoutes(server *gin.Engine, loginCheck gin.HandlerFunc) {
	rg := server.Group("/v2/article/")
	rg.Use(loginCheck)

	rg.POST("/getArticles", r.GetArticles)
	rg.POST("/getAuthorArticles", r.GetAuthorArticles)
}

func (r *ArticleHandler) GetArticles(c *gin.Context) {
	req := domain.QueryArticlesReq{}
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
	data, err := r.svc.GetArticles(c, req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
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
