package middleware

import (
	"github.com/basicprojectv2/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type LoginJWTMiddlewareBuilder struct {
}

// 校验jwt token
func (m *LoginJWTMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		// 如果是登录/注册/短信登录/发送短信 这几个接口不需要校验jwt token
		if path == "/v2/users/loginin" ||
			path == "/v2/users/signin" ||
			path == "/v2/users/loginsms/code/send" ||
			path == "/v2/users/loginsms" {
			return
		}
		// 获取jwt token
		authCode := ctx.GetHeader("Authorization")
		if authCode == "" {
			// token为空
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 解析jwt token
		parts := strings.SplitN(authCode, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 如果校验无误 设置id到上下文当中
		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("userid", claims.UserId)
	}
}
