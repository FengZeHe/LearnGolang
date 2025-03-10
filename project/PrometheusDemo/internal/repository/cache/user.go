package cache

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"prometheusdemo/internal/domain"
	"time"
)

type UserCache interface {
	Get(ctx context.Context, key string) ([]domain.User, error)
	Set(ctx context.Context, key, data string) error
}

type RedisUserCache struct {
	cmd        redis.Cmdable
	expiration time.Duration
}

func NewUserCache(cmd redis.Cmdable) UserCache {
	return &RedisUserCache{
		cmd:        cmd,
		expiration: time.Minute * 15,
	}
}

func (r RedisUserCache) Get(ctx context.Context, key string) (users []domain.User, err error) {
	data, err := r.cmd.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(data), &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r RedisUserCache) Set(ctx context.Context, key, data string) (err error) {
	if err = r.cmd.Set(ctx, key, data, r.expiration).Err(); err != nil {
		return err
	}
	return nil
}
