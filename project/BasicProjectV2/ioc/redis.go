package ioc

import (
	"github.com/basicprojectv2/settings"
	"github.com/redis/go-redis/v9"
)

func InitRedis(Conf *settings.RedisConfig) (rdb redis.Cmdable) {
	rdb = redis.NewClient(&redis.Options{
		Addr: Conf.Host,
	})
	return rdb
}
