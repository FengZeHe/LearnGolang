package cache

import (
	"BasicProject/models"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

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

// 用于用户短信验证码登录使用，存取手机号和验证码
func SetCodeForUserSMSLogin(phone, code string) (err error) {
	key := fmt.Sprintf("%s%s", KeyUserSMSLoginSet, phone)
	if err = rdb.Set(ctx, key, code, 5*time.Minute).Err(); err != nil {
		log.Println("set SMS data ERROR:", err)
		return err
	}
	return nil
}
