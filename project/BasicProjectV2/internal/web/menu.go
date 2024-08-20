package web

import (
	"github.com/basicprojectv2/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type MenuHandler struct {
	svc service.MenuService
}

func NewMenuHandler(svc service.MenuService) *MenuHandler {
	return &MenuHandler{svc: svc}
}

// 注册路由
func (h *MenuHandler) RegisterRoutes(server *gin.Engine, loginCheck gin.HandlerFunc) {
	ug := server.Group("/v2/menus/")
	// 获取菜单
	ug.GET("/list", loginCheck, h.GetMenus)
	// 创建菜单
	ug.POST("")
	// 更新菜单
	ug.PUT("")
	// 删除菜单
	ug.DELETE("")

}

func (h *MenuHandler) GetMenus(ctx *gin.Context) {
	menu, err := h.svc.GetMenus(ctx)
	if err != nil {
		log.Println("错误", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "系统错误",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": menu,
	})

}
