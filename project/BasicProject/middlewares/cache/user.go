package cache

import (
	"BasicProject/models"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"log"
	"strconv"
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
			return key, false, nil
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

// 用于短信验证码剩余次数
func CheckSMSResidualDegree(phone string) (status bool, err error) {
	/*
		1.首先查询 key是否存在，如果存在就不能覆盖
		2.key如果存在，则查询value剩余次数并 -1
		3.key如果不存在则添加一个 并设置剩余次数5次， 过期时间600s
		查询该手机发送验证码的的剩余次数，如果 -1次大于0，那么可以发送 return true
	*/
	tempKey := fmt.Sprintf("%s%s", KeyUserSMSCount, phone)
	if existed := ExistKey(tempKey); existed != true {
		// key不存在 设置值
		if err = rdb.Set(ctx, tempKey, 5, time.Second*600).Err(); err != nil {
			// 设置成功
			return true, nil
		} else {
			return false, err
		}
	} else {
		// key 存在 查询值
		res, err := rdb.Get(ctx, tempKey).Result()
		if err != nil {
			// 查询的时候出错 返回错误
			return false, err
		}
		intres, _ := strconv.Atoi(res)
		if intres-1 > 0 {
			// 还有剩余次数 -1 并返回true
			if err = rdb.Decr(ctx, tempKey).Err(); err != nil {
				// 操作redis错误
				return false, err
			}
			return true, nil
		}

	}

	return true, nil
}
