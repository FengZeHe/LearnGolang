package web

import (
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DraftHandler struct {
	svc service.DraftService
}

func NewDraftHandler(svc service.DraftService) *DraftHandler {
	return &DraftHandler{svc: svc}
}

func (r *DraftHandler) RegisterRoutes(server *gin.Engine, loginCheck gin.HandlerFunc) {
	rg := server.Group("/v2/draft/")
	rg.Use(loginCheck)
	rg.GET("/getArticles", r.GetArticles)
	rg.POST("/addArticle", r.AddArticle)
	rg.POST("/updateArticle", r.AddArticle)
}

func (r *DraftHandler) GetArticles(c *gin.Context) {

}

func (r *DraftHandler) AddArticle(c *gin.Context) {
	var req domain.AddDraftReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}
	userid, exists := c.Get("userid")
	if !exists {
		c.JSON(400, gin.H{
			"msg": "用户未登录",
		})
		return
	}
	useridStr := userid.(string)
	if err := r.svc.AddArticle(c, req, useridStr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "保存成功",
	})
}

func (r *DraftHandler) UpdateArticle(c *gin.Context) {

}
