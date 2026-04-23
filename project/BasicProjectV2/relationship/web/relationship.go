package web

import (
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
}
