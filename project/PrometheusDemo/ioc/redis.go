package ioc

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"prometheusdemo/internal/web/middlewares"
)

func InitRedis(config *RedisConfig) (rdb redis.Cmdable) {
	rdb = redis.NewClient(&redis.Options{
		Addr: config.Addr,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Println("redis connect error:", err)
		panic(err)
	}
	log.Println("redis connect success")

	return rdb
}

func InitRedisClient(config *RedisConfig) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr: config.Addr,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Println("redis connect error:", err)
		panic(err)
	}
	log.Println("redis client success")
	// Add Hook
	hook := &middlewares.CacheHitHook{
		NameSpace: "my_http",
		Subsystem: "prometheus_demo",
		Name:      "cache_hit",
		Help:      "xxx",
	}
	client.AddHook(hook)

	return client
}
