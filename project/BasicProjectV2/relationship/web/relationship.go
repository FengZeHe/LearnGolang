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
