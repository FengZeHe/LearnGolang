package localcache

import "github.com/coocood/freecache"

// 设置本地缓存大小
var (
	cacheSize  = 100 * 1024 * 1024
	LocalCache = freecache.NewCache(cacheSize)
)

// 设置缓存
func SetCache(key, value string, expire int) (err error) {
	err = LocalCache.Set([]byte(key), []byte(value), expire)
	if err != nil {
		return err
	}
	return nil
}

// 获取缓存
func GetCache(key string) (value []byte, err error) {
	value, err = LocalCache.Get([]byte(key))
	if err != nil {
		return nil, err
	}
	return value, nil
}

// 删除缓存
func DelCache(key string) (affected bool) {
	affected = LocalCache.Del([]byte(key))
	return affected
}
