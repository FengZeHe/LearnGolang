package localcache

import (
	"BasicProject/middlewares/cache"
	"BasicProject/models"
	"encoding/json"
	"fmt"
	"log"
)

// 根据userid设置localcache
func SetLocalCacheByUserId(data *models.User, userid string) (err error) {
	strdata, _ := json.Marshal(data)
	key := fmt.Sprintf("%s%s", cache.KeyUserIdSet, userid)
	if err = SetCache(key, string(strdata), 600); err != nil {
		log.Println("Set Local Cache ERROR", err)
		return err
	}
	log.Println("Set Local Cache success!")
	return nil
}

// 通过userid获取缓存
func GetLocalCacheByUserId(key string) (data *models.User, err error) {
	strKey := fmt.Sprintf("%s%s", cache.KeyUserIdSet, key)
	res, err := GetCache(strKey)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(res, &data)
	return data, nil
}

// 通过userid删除缓存
func DelLocalCacheByUserId(key string) (err error) {
	strKey := fmt.Sprintf("%s%s", cache.KeyUserIdSet, key)
	affected := DelCache(strKey)
	if affected != true {
		// 删除失败？
		log.Println("没有删除条目")
	}
	return nil
}
