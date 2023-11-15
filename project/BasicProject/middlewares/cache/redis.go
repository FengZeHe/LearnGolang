package cache

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

// redis初始化连接
func RedisInit() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	log.Println("Redis init success!")
	return nil
}

// 判断key是否存在
func ExistKey(key string) bool {
	n, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		log.Println("Get Exist key Error", err)
		return false
	}
	if n == 0 {
		log.Println("key not exist")
		return false
	}
	log.Println("key exist")
	return true
}

// 缓存的中间件
func CacheMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 首先要获取上下文拿到email,先判断 Key是否存在 ，如果存在则redis中查找，如果不存在则 c.Next
		userid, _ := c.Get("userid")
		if existed := ExistKey(fmt.Sprintf("%v", userid)); existed == true {
			// 在redis中查找
		} else {
			c.Next()
		}
	}
}
