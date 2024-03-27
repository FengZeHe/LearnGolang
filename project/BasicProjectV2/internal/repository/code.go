package repository

import (
	"context"
	"github.com/basicprojectv2/internal/repository/cache"
)

type CodeRepository interface {
	SetCode(ctx context.Context, biz, phone, code string) error
	VerifyCode(ctx context.Context, biz, phone, code string) (bool, error)
}

type CachedCodeRepository struct {
	cache cache.CodeCache
}

func NewCodeRepository(c cache.CodeCache) CodeRepository {
	return &CachedCodeRepository{
		cache: c,
	}
}

func (c *CachedCodeRepository) SetCode(ctx context.Context, biz, phone, code string) (err error) {
	return c.cache.SetCode(ctx, biz, phone, code)
}

func (c *CachedCodeRepository) VerifyCode(ctx context.Context, biz, phone, code string) (result bool, err error) {
	return c.cache.VerifyCode(ctx, biz, phone, code)
}

// 执行lua脚本的函数
func EvalLuaScript(key string, value string, luaScript string) {

}
