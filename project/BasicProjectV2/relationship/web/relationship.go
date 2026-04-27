package web

import (
	"net/http"

	"github.com/basicprojectv2/relationship/domain"
	"github.com/basicprojectv2/relationship/service"
	"github.com/gin-gonic/gin"
)

type RelationshipHandler struct {
	svc service.RelationshipService
}

func NewRelationshipHandler(svc service.RelationshipService) *RelationshipHandler {
	return &RelationshipHandler{svc: svc}
}

func (r *RelationshipHandler) RegisterRoutes(service *gin.Engine, loginCheck gin.HandlerFunc) {
	rg := service.Group("/v2/relationship/")
	rg.Use(loginCheck)
	rg.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong rrr",
		})
	})
	rg.POST("/follow", r.HandleFollow)
	rg.POST("/block", r.HandleBlock)
	rg.GET("/", r.QueryRelationship)
	rg.GET("/followee", r.QueryFolloweeList)
	rg.GET("/follower", r.QueryFollowerList)
	rg.GET("/count", r.HandleCount)
	rg.GET("/countMe", r.HandleCountMe)
}

func (r *RelationshipHandler) HandleFollow(ctx *gin.Context) {
	uid, exists := ctx.Get("userid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}
	strUid, _ := uid.(string)

	req := domain.FollowReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := r.svc.HandleFollow(strUid, req, ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "ok",
	})

}

func (r *RelationshipHandler) HandleBlock(ctx *gin.Context) {
	uid, exists := ctx.Get("userid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}
	strUid, _ := uid.(string)
	req := domain.BlockReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := r.svc.HandleBlock(strUid, req, ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "ok",
	})

}

func (r *RelationshipHandler) QueryRelationship(ctx *gin.Context) {
	uid, exists := ctx.Get("userid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}
	strUid, _ := uid.(string)

	targetUid := ctx.Query("uid")

	userStatus, _ := r.svc.QueryRelationship(strUid, targetUid, ctx)
	ctx.JSON(200, gin.H{
		"data": userStatus,
	})

}

func (r *RelationshipHandler) QueryFolloweeList(ctx *gin.Context) {
	uid, exists := ctx.Get("userid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}
	strUid, _ := uid.(string)
	req := domain.FollowListReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		req.PageIndex = 1
		req.PageSize = 10
	}

	fl, err := r.svc.QueryFolloweeList(strUid, req, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": fl,
	})
}

func (r *RelationshipHandler) QueryFollowerList(ctx *gin.Context) {
	uid, exists := ctx.Get("userid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}
	strUid, _ := uid.(string)
	req := domain.FollowListReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		req.PageIndex = 1
		req.PageSize = 10
	}

	fl, err := r.svc.QueryFollowerList(strUid, req, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": fl,
	})
}

func (r *RelationshipHandler) HandleCount(ctx *gin.Context) {
	uid := ctx.Query("uid")
	c, err := r.svc.CountRelationship(uid, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": c,
	})
}

func (r *RelationshipHandler) HandleCountMe(ctx *gin.Context) {
	uid, exists := ctx.Get("userid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}

	strUid, _ := uid.(string)
	c, err := r.svc.CountRelationship(strUid, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": c,
	})
}
