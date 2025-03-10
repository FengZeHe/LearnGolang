package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"prometheusdemo/internal/repository"
)

type MiddleUserCache struct {
	userRepo repository.UserRepository
	rdb      *redis.Client
}

func NewMiddleUserCache(userRepo repository.UserRepository, rdb *redis.Client) *MiddleUserCache {
	return &MiddleUserCache{userRepo: userRepo, rdb: rdb}
}

func (cache *MiddleUserCache) GetUserCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("pass middleware")
		data := cache.rdb.Get(context.Background(), "users")
		log.Println("rdb GET", data)
		c.Next()
		return
	}
}
