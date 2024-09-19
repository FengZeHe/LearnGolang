package web

import (
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

	rg.GET("", r.GetArticles)
}

func (r *ArticleHandler) GetArticles(c *gin.Context) {
	_, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户未登录",
		})
		return
	}

	data, err := r.svc.GetArticles(c)
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
