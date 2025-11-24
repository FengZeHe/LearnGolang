package web

import (
	"context"
	"net/http"

	"github.com/basicprojectv2/interactive/domain"
	"github.com/basicprojectv2/interactive/service"
	"github.com/gin-gonic/gin"
)

type InteractiveHandler struct {
	svc service.InteractiveService
}

func NewInteractiveHandler(svc service.InteractiveService) *InteractiveHandler {
	return &InteractiveHandler{
		svc: svc,
	}
}

func (r *InteractiveHandler) RegisterRoutes(server *gin.Engine, loginCheck gin.HandlerFunc) {
	rg := server.Group("/v2/interactive/")
	rg.Use(loginCheck)
	rg.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	rg.POST("/addRead", r.AddReadCount)
	rg.POST("/like", r.HandleLike)
	rg.POST("/collect", r.HandleCollect)
}

/*
增加文章阅读数统计
*/
func (r *InteractiveHandler) AddReadCount(c *gin.Context) {
	req := domain.AddReadCountReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
		return
	}
	if err := r.svc.AddReadCount(req.Aid, context.Background()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "add count error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "add count success",
	})
}

// 处理用户点赞
func (r *InteractiveHandler) HandleLike(c *gin.Context) {

}

// 处理用户收藏
func (r *InteractiveHandler) HandleCollect(c *gin.Context) {}
