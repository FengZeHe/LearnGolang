package web

import (
	"net/http"

	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/service"
	"github.com/gin-gonic/gin"
)

type UserSettingHandler struct {
	usService service.UserSettingService
}

func NewUserSettingHandler(usService service.UserSettingService) *UserSettingHandler {
	return &UserSettingHandler{usService: usService}
}

// 注册路由
func (h *UserSettingHandler) RegisterRoutes(server *gin.Engine, loginCheck gin.HandlerFunc) {
	ug := server.Group("/v2/usersetting/")
	ug.Use(loginCheck)

	ug.GET("/", h.GetUserSetting)
	ug.POST("/", h.HandleUserSetting)
	ug.POST("/reset", h.ResetUserSetting)

}

// 获取用户配置
func (h *UserSettingHandler) GetUserSetting(ctx *gin.Context) {
	uid, _ := ctx.Get("userid")
	var us = domain.UserSetting{}
	var err error
	if us, err = h.usService.GetUserSetting(uid.(string), ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, us)
}

// 更新用户配置
func (h *UserSettingHandler) HandleUserSetting(ctx *gin.Context) {
	req := domain.UserSettingReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, _ := ctx.Get("userid")
	if err := h.usService.HandleUserSetting(uid.(string), req, ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

// 重置用户配置
func (h *UserSettingHandler) ResetUserSetting(ctx *gin.Context) {

}
