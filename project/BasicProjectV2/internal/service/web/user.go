package web

import (
	"github.com/basicprojectv2/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.POST("/hi", h.Hi)
}

func (h *UserHandler) Hi(ctx *gin.Context) {
	ctx.JSON(200, "msg:hello")
}
