package cache

import (
	"BasicProject/setting"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

// redis初始化连接
func RedisInit(cfg *setting.Redis) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
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
		//log.Println("key not exist")
		return false
	}
	//log.Println("key exist")
	return true
}

// 缓存的中间件
func RedisCacheMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 首先要获取上下文拿到email,先判断 Key是否存在 ，如果存在则redis中查找，如果不存在则 c.Next
		userid, _ := c.Get("userid")
		key := fmt.Sprintf("%s%s", KeyUserIdSet, userid)
		if existed := ExistKey(fmt.Sprintf("%v", key)); existed == true {
			// 在redis中查找
			res, err := GetCacheByUseId(key)
			if err != nil {
				log.Println("Get Cache Error", err)
				c.Next()
			}
			log.Println("Hit Redis Cache ")
			c.JSON(http.StatusOK, gin.H{
				"message":     "success",
				"userprofile": res,
			})
			c.Abort()
		} else {
			log.Println("Miss Redis Cache")
			c.Next()
		}
	}
}

// 删除某个key
func DeleteKey(key string) (err error) {
	if err = rdb.Del(ctx, key).Err(); err != nil {
		log.Println("redis Delete Key ERROR", err)
		return err
	}

	return nil
}

// 执行lua脚本的函数
func EvalLuaScript(key string, value string, luaScript string) (interface{}, error) {
	result, err := rdb.Eval(ctx, luaScript, []string{key}, value).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}
