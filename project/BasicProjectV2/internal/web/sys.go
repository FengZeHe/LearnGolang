package web

import (
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/service"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type SysHandler struct {
	svc      service.SysService
	enforcer *casbin.Enforcer
}

func NewSysHandler(svc service.SysService, enforcer *casbin.Enforcer) *SysHandler {
	return &SysHandler{
		svc:      svc,
		enforcer: enforcer,
	}
}

// 注册路由
func (h *SysHandler) RegisterRoutes(server *gin.Engine, roleCheck, loginCheck gin.HandlerFunc) {
	ug := server.Group("/v2/sys/")
	ug.GET("/menu", loginCheck, h.HandleUserGetMenu)
	ug.GET("/hi", loginCheck, roleCheck, h.Hi)
	ug.GET("/menuList", loginCheck, h.HandleGetMenu)
	ug.GET("/roleList", loginCheck, h.HandleGetRole)
	ug.GET("/apiList", loginCheck, h.HandleGetAPI)

	// 管理casbin策略
	ug.POST("/addPolicy", loginCheck, h.HandleAddPolicy)
	ug.POST("/updatePolicy", loginCheck, h.HandleUpdatePolicy)
	ug.POST("/deletePolicy", loginCheck, h.HandleDeletePolicy)

}

func (h *SysHandler) Hi(ctx *gin.Context) {
	ctx.JSON(200, "Hi!!!")
}

// 添加Casbin策略
func (h *SysHandler) HandleAddPolicy(ctx *gin.Context) {
	var req domain.AddCasbinRulePolicyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
	}
	ok, err := h.enforcer.AddPolicy(req.NewPolicy)
	if err != nil {
		return
	}
	if ok {
		log.Println("add policy success")
		ctx.JSON(http.StatusOK, "success")
	} else {
		log.Println("add policy failed")
		ctx.JSON(http.StatusInternalServerError, "failed")
	}

}

// todo  更新casbin策略
func (h *SysHandler) HandleUpdatePolicy(ctx *gin.Context) {
	//先删除，再添加
	var req domain.UpdateCasbinPolicyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		return
	}

	ok, err := h.enforcer.RemovePolicy(req.OldPolicy)
	if err != nil {
		log.Println("remove policy fail", err)
		return
	}
	ok, err = h.enforcer.AddPolicy(req.NewPolicy)
	if err != nil {
		log.Println("add policy fail", err)
		return
	}

	if ok {
		log.Println("update policy success")
		ctx.JSON(http.StatusOK, "success")
	}

}

// 删除casbin策略
func (h *SysHandler) HandleDeletePolicy(ctx *gin.Context) {
	var req domain.RemoveCasbinPolicyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
	}
	ok, err := h.enforcer.RemovePolicy(req.RemovePolicy)
	if err != nil {
		return
	}
	if ok {
		log.Println("remove policy success")
		ctx.JSON(http.StatusOK, "success")

	} else {
		log.Println("remove policy failed")
		ctx.JSON(http.StatusInternalServerError, "failed")

	}

}

func (h *SysHandler) HandleGetAPI(ctx *gin.Context) {
	al, err := h.svc.GetAPI(ctx)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": al,
	})
}

func (h *SysHandler) HandleGetRole(ctx *gin.Context) {
	rl, err := h.svc.GetRole(ctx)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": rl,
	})
}

// 返回菜单目录
func (h *SysHandler) HandleGetMenu(ctx *gin.Context) {
	m, err := h.svc.GetMenu(ctx)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": m,
	})
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
	menus, err := h.svc.GetMenuByUserID(ctx, strUserid)
	if err != nil {
		log.Println("svc GetMenuByUserID err:", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": menus,
	})
}
