package localcache

import (
	"github.com/coocood/freecache"
	"github.com/gin-gonic/gin"
)

var (
	cacheSize  = 100 * 1024 * 1024
	LocalCache = freecache.NewCache(cacheSize)
)

// 本次缓存中间件
func LocalCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func SetCache(key, value string, expire int) (err error) {
	err = LocalCache.Set([]byte(key), []byte(value), expire)
	if err != nil {
		return err
	}
	return nil
}

func GetCache(key string) (value []byte, err error) {
	value, err = LocalCache.Get([]byte(key))
	if err != nil {
		return nil, err
	}
	return value, nil
}
