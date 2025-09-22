package web

import (
	"net/http"

	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/service"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	svc service.CommentService
}

func NewCommentHandler(svc service.CommentService) *CommentHandler {
	return &CommentHandler{svc: svc}
}

func (h *CommentHandler) RegisterRoutes(server *gin.Engine, loginCheck gin.HandlerFunc) {
	cg := server.Group("/v2/comment/")
	cg.POST("/add", loginCheck, h.AddComment)
}

func (h *CommentHandler) AddComment(c *gin.Context) {
	var req domain.AddCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
