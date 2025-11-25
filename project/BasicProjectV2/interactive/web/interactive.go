package web

import (
	"context"
	"log"
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

	rg.GET("/status/:aid", r.GetStatus)
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
	req := domain.LikeReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
		return
	}

	uid, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "unauthorized",
		})
		return
	}
	uidStr := uid.(string)

	if err := r.svc.HandleLike(req.Aid, req.Like, uidStr, context.Background()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "like error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "like success",
	})

}

// 处理用户收藏
func (r *InteractiveHandler) HandleCollect(c *gin.Context) {
	req := domain.CollectReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
	}

	uid, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "unauthorized",
		})
		return
	}
	uidStr := uid.(string)

	if err := r.svc.HandleCollect(req.Aid, req.Collect, uidStr, context.Background()); err != nil {
		log.Println("error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "collect error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "collect success",
	})

}

// 返回用户与文章交互
func (r *InteractiveHandler) GetStatus(c *gin.Context) {
	aid := c.Param("aid")
	c.JSON(http.StatusOK, gin.H{
		"aid": aid,
	})

	uid, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "unauthorized",
		})
		return
	}
	uidStr := uid.(string)

	res, err := r.svc.GetStatus(aid, uidStr, context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "get status error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})

}
