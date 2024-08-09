package web

import (
	"github.com/basicprojectv2/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
func (h *SysHandler) RegisterRoutes(server *gin.Engine, roleCheck, loginCheck gin.HandlerFunc) {
	ug := server.Group("/v2/sys/")
	ug.GET("/menu", loginCheck, h.HandleUserGetMenu)
	ug.GET("/hi", loginCheck, roleCheck, h.Hi)

}

func (h *SysHandler) Hi(ctx *gin.Context) {
	ctx.JSON(200, "Hi!!!")
}

// 处理获取菜单请求
func (h *SysHandler) HandleUserGetMenu(ctx *gin.Context) {
	userid, exists := ctx.Get("userid")
	if !exists {
		ctx.JSON(400, gin.H{
			"msg": "用户未登录",
		})
	}
	strUserid := userid.(string)
	menus, err := h.svc.GetMenu(ctx, strUserid)
	if err != nil {
		log.Println("svc GetMenu err:", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": menus,
	})
}
