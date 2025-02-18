package web

import (
	"github.com/basicprojectv2/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleHandler struct {
	svc service.RoleService
}

func NewRoleHandler(svc service.RoleService) *RoleHandler {
	return &RoleHandler{svc: svc}
}

func (r *RoleHandler) RegisterRoutes(server *gin.Engine, loginCheck gin.HandlerFunc) {
	rg := server.Group("/v2/roles/")

	rg.Use(loginCheck)
	rg.GET("/list", r.HandleGetRoles)
	rg.POST("/updateRoles")
}

// todo 批量更新规则
func (r *RoleHandler) HandleUpdateRoles(c *gin.Context) {

}

func (r *RoleHandler) HandleGetRoles(ctx *gin.Context) {
	roles, err := r.svc.GetRoles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "系统错误",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": roles,
	})
}
