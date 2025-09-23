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
	cg.GET("/", loginCheck, h.GetComment)
	cg.DELETE("/", loginCheck, h.DeleteComment)
}

func (h *CommentHandler) GetComment(c *gin.Context) {
	aid := c.Query("aid")
	comments, err := h.svc.GetComment(c, aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}

func (h *CommentHandler) AddComment(c *gin.Context) {
	var req domain.AddCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.AddComment(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "comment added"})
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	id := c.Query("id")
	if err := h.svc.DeleteComment(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "comment deleted"})
}
