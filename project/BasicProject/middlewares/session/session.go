package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user-session")
		if user == nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "没有登录"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func InitSession(r *gin.Engine) {
	//初始化cookie存储引擎，还有其他不同的存储引擎如redis...
	//"my_secret"是用于加密的密钥
	store := cookie.NewStore([]byte("my_secret"))
	r.Use(sessions.Sessions("user-session", store))
}
