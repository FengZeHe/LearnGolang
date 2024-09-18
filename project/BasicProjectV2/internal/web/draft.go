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
	rg.POST("/getDraft", r.getDraft)
	rg.POST("/addArticle", r.AddArticle)
	rg.POST("/updateArticle", r.UpdateArticle)
	rg.POST("/deleteArticle", r.DeleteArticle)
}

func (r *DraftHandler) GetArticles(c *gin.Context) {
	userid, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户未登录",
		})
	}
	useridStr := userid.(string)
	articles, err := r.svc.GetArticles(c, useridStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": articles,
	})

}

func (r *DraftHandler) getDraft(c *gin.Context) {
	var req domain.GetDraftReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}
	draft, err := r.svc.GetDraft(c, req.DraftID, req.AuthorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": draft,
	})
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
	var req domain.UpdateDraftReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.svc.UpdateArticle(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
	})

}

func (r *DraftHandler) DeleteArticle(c *gin.Context) {
	var req domain.DeleteDraftReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.svc.DeleteArticle(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})

}
