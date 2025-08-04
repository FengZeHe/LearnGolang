package limiter

import (
	"context"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type TokenBucketLimiter struct {
	limiter *rate.Limiter
}

func NewTokenBucketLimiter(r rate.Limit, b int) *TokenBucketLimiter {
	return &TokenBucketLimiter{limiter: rate.NewLimiter(r, b)}
}

func (t *TokenBucketLimiter) LimitInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if !t.limiter.Allow() {
			log.Println("请求过于频繁")
			return nil, status.Errorf(codes.ResourceExhausted, "请求过于频繁，请稍后再试")
		}
		return handler(ctx, req)
	}
}
