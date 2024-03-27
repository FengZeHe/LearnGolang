package cache

import (
	"context"
	_ "embed"
	//"errors"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"log"
)

// 引入lua脚本
var (
	//go:embed lua/setCode.lua
	luaSetCode string
	//go:embed lua/smsGetDegree.lua
	smsGetDegree string
	//go:embed lua/smsVerify.lua
	smsVerify string
)

type RedisCodeCache struct {
	cmd redis.Cmdable
}

type CodeCache interface {
	SetCode(ctx context.Context, biz, phone, code string) (err error)
	VerifyCode(ctx context.Context, biz, phone, code string) (result bool, err error)
}

func NewCodeCache(cmd redis.Cmdable) CodeCache {
	return &RedisCodeCache{
		cmd: cmd,
	}
}

func (c *RedisCodeCache) SetCode(ctx context.Context, biz, phone, code string) (err error) {
	//执行lua脚本
	// 判断是否能发送验证码
	/*
		1. 查询该手机号发送验证码的剩余次数
		2. 如果剩余此时 > 0 则允许发送
		3.
	*/
	res, err := c.cmd.Eval(ctx, smsGetDegree, []string{phone}, code).Int()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return
		}
		return err
	}
	switch res {
	case -1:
		return errors.New("请求发送过于频繁")
	case 1:
		log.Println("发送验证码")
	}

	// 发送验证码
	/*
		1. 发送验证码，设置剩余次数 -1
		2. 记录发送的验证码到redis中
	*/
	res, err = c.cmd.Eval(ctx, luaSetCode, []string{phone}, code).Int()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		} else {
			return err
		}
	}
	switch res {
	case 1:
		return nil
	}
	return nil
}

func (c *RedisCodeCache) VerifyCode(ctx context.Context, biz, phone, code string) (result bool, err error) {
	// 执行lua脚本
	res, err := c.cmd.Eval(ctx, smsVerify, []string{phone}, code).Int()
	if err != nil {
		log.Println("err=>", err)
		result = false
		return result, err
	}
	switch res {
	case -1:
		result = false
	case 1:
		result = true
	}

	return result, nil
}

func (c *RedisCodeCache) key(biz, phone string) string {
	return fmt.Sprintf("phone_code:%s:%s", biz, phone)
}
