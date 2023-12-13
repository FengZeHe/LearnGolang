package cache

import (
	"BasicProject/models"
	"BasicProject/setting"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"time"
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
func CacheMiddleWare(kind string) gin.HandlerFunc {
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
			c.JSON(http.StatusOK, gin.H{
				"message":     "success",
				"userprofile": res,
			})
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// 根据用户id获取缓存
func GetCacheByUseId(key string) (data *models.User, err error) {
	res, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Println("GET redis CACHE ERROR", err)
		return nil, err
	}
	err = json.Unmarshal([]byte(res), &data)
	return data, nil
}

// 根据用户id设置缓存
func SetCacheByUserId(data *models.User, userid string) (err error) {
	strdata, _ := json.Marshal(data)
	key := fmt.Sprintf("%s%s", KeyUserIdSet, userid)
	if err = rdb.Set(ctx, key, strdata, 60*time.Second).Err(); err != nil {
		log.Println("set Cache ERROR:", err)
		return err
	}
	log.Println("set cache success!")
	return nil
}
