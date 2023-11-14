package cache

import (
	"context"
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
