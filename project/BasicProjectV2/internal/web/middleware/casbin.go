package middleware

import (
	"github.com/basicprojectv2/internal/repository"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CasbinRoleCheck struct {
	Enforcer *casbin.Enforcer
	UserRepo repository.UserRepository
}

func NewCasbinRoleCheck(enforcer *casbin.Enforcer, UserRepo repository.UserRepository) *CasbinRoleCheck {
	return &CasbinRoleCheck{Enforcer: enforcer, UserRepo: UserRepo}
}

func (c *CasbinRoleCheck) CheckRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// todo  在这里进行casbin role 的验证
		userid, exists := ctx.Get("userid")
		if !exists {
			ctx.Abort()
		}
		userIDStr := userid.(string)
		// todo 通过userID查询role
		user, err := c.UserRepo.FindById(ctx, userIDStr)
		if err != nil {
			ctx.Abort()
		}
		methods := ctx.Request.Method
		url := ctx.Request.URL.String()
		log.Println(user.Email, url, methods)

		err = c.Enforcer.LoadPolicy()
		if err != nil {
			log.Println("Enforcer load policy error", err)
		}

		// 检查权限
		ok, err := c.Enforcer.Enforce(user.Email, url, methods)
		if err != nil {
			log.Println("Enforce failed", err)
			ctx.Abort()
		}
		if ok != true {
			ctx.AbortWithStatus(http.StatusBadRequest)
		} else {
			ctx.Next()
		}
	}
}
