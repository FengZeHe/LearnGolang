package web

import (
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

	rg.GET("/hi", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi",
		})
	})
}
