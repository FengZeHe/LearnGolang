package web

import (
	"github.com/gin-gonic/gin"
	"prometheusdemo/internal/service"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/user")
	ug.POST("/getUser")
}

func (h *UserHandler) UserHandler(c *gin.Context) {

}
