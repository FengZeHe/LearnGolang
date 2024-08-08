package web

import (
	"github.com/basicprojectv2/internal/service"
	"github.com/basicprojectv2/internal/web/middleware"
	"github.com/gin-gonic/gin"
)

type SysHandler struct {
	svc service.SysService
}

func NewSysHandler(svc service.SysService) *SysHandler {
	return &SysHandler{
		svc: svc,
	}
}

// 注册路由
func (h *SysHandler) RegisterRoutes(server *gin.Engine) {
	m := middleware.LoginJWTMiddlewareBuilder{}
	ug := server.Group("/v2/sys/")
	ug.GET("/menu", m.CheckLogin(), h.Hi)
}

func (h *SysHandler) Hi(ctx *gin.Context) {
	ctx.JSON(200, "hello Menu")
}
