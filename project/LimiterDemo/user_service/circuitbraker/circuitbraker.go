package circuitbraker

import (
	"context"
	"github.com/go-kratos/aegis/circuitbreaker"
	"github.com/go-kratos/aegis/circuitbreaker/sre"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strings"
	"sync/atomic"
)

type CircuitBreaker struct {
	cb      circuitbreaker.CircuitBreaker // aegis熔断器
	enabled int32                         //1-禁用熔断器  0-熔断器闭合(正常逻辑)
	manual  int32                         //0-不使用手动开关 1-强制闭合，全部请求通过  2-强制打开，全部请求拒绝
}

func NewCircuitBraker(opts ...sre.Option) *CircuitBreaker {
	return &CircuitBreaker{
		cb:      sre.NewBreaker(opts...),
		enabled: 0, // 默认闭合 正常逻辑
		manual:  0, // 默认不使用手动开关
	}
}

// 手动关闭熔断器，允许所有请求
func (c *CircuitBreaker) ManualClose() {
	atomic.StoreInt32(&c.manual, 1)
}

// 手动打开熔断器，拦截所有请求
func (c *CircuitBreaker) ManualOpen() {
	atomic.StoreInt32(&c.manual, 2)
}

// 不适用手动控制
func (c *CircuitBreaker) DisableManul() {
	atomic.StoreInt32(&c.manual, 0)
}

// 熔断器闭合-正常逻辑
func (c *CircuitBreaker) Enable() {
	atomic.StoreInt32(&c.enabled, 0)
}

// 禁用熔断器-允许所有请求
func (c *CircuitBreaker) Disable() {
	atomic.StoreInt32(&c.enabled, 1)
}

// 检查熔断器状态
func (c *CircuitBreaker) IsEnabled() int32 {
	return atomic.LoadInt32(&c.enabled)
}

// 检查熔断器手动状态
func (c *CircuitBreaker) IsManual() int32 {
	return atomic.LoadInt32(&c.manual)
}

func (c *CircuitBreaker) CircuitBrakerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 如果是控制接口则直接通过
		if info.FullMethod == "/user_service.UserService/ControlCircuitBraker" {
			log.Println("控制接口 直接通过")
			return handler(ctx, req)
		}

		switch c.manual {
		case 0: // 正常逻辑
			if atomic.LoadInt32(&c.enabled) == 0 { // 走正常逻辑判断
				if err := c.cb.Allow(); err != nil {
					log.Println("自动触发熔断，拒绝请求", err)
					return nil, status.Errorf(codes.Unavailable, "熔断器自动触发熔断")
				}
				resp, err = handler(ctx, req)
				if err != nil {
					c.cb.MarkFailed()
					log.Println("请求失败，上报", err)
				} else {
					c.cb.MarkSuccess()
					log.Println("请求成功，上报")
				}
			} else if atomic.LoadInt32(&c.enabled) == 1 { // 禁用熔断器，允许所有请求
				log.Println("已禁用拦截器，允许所有请求")
				return handler(ctx, req)
			}
		case 1: // 手动闭合熔断器，允许所有请求
			log.Println("手动闭合熔断器，允许所有请求")
			return handler(ctx, req)
		case 2: // 手动打开熔断器，拒绝所有请求

			/*
				降级部分：
				1.可以使用全路径
				2.可以匹配前缀
			*/
			if strings.HasPrefix(info.FullMethod, "/user_service.UserService/CoreBusiness") {
				log.Println("核心业务 通过")
				return handler(ctx, req)
			}

			log.Println("手动打开熔断器，拒绝所有请求")
			return nil, status.Errorf(codes.Unavailable, "熔断器已手动强制打开")
		}
		return resp, nil
	}
}
