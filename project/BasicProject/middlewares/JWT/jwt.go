package JWT

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"time"
)

// 定义加密的key
var jwtKey = []byte("my_secret")

// MyClaims 定义jwt对象struct
type MyClaims struct {
	UserId string `json:"userid"`
	jwt.RegisteredClaims
}

// 定义jwt中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			1. 检查token是否为空
			2. 检查token的格式是否正确
			3. 分割token并解析是否正确
		*/

		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": "请求未携带token, 无法访问",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(token, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"message": "token请求格式有误",
			})
			c.Abort()
			return
		}
		claims, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "无效的token",
			})
			c.Abort()
			return
		}
		//将用户id 设置到上下文
		c.Set("userid", claims.UserId)
		c.Next()
	}
}

// 解析token
func ParseToken(tokenStr string) (myclaims *MyClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if myclaims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return myclaims, nil
	}
	return nil, errors.New("invalid token")
}

// 生成token
func GenToken(userId string) (tokenStr string, err error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &MyClaims{UserId: userId, RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
