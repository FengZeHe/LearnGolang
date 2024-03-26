package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
)

var jwtKey = []byte("my_secret")

type UserClaims struct {
	jwt.RegisteredClaims
	UserId string
}

// 生成token
func GenToken(userId string) (tokenStr string, err error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &UserClaims{UserId: userId, RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// 解析token
func ParseToken(tokenStr string) (myclaims *UserClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if myclaims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return myclaims, nil
	}
	return nil, errors.New("invalid token")
}
