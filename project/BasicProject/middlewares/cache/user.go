package cache

import (
	"BasicProject/models"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
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

// 用于验证短信登录，手机号和验证码是否正确
func VerifyCodeForUserSMSLogin(phone, code string) (key string, res bool, err error) {
	key = fmt.Sprintf("%s%s", KeyUserSMSLoginSet, phone)
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return key, true, nil
		}
		fmt.Println("Get Redis Result ERROR", err)
		return key, false, err
	}
	if val != code {
		fmt.Println("Code ERROR 验证码不正确，登录失败")
		return key, false, nil
	}
	return key, true, nil
}
