package main

import (
	"context"
	"github.com/go-kratos/aegis/circuitbreaker/sre"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	cb "limiterdemo/user_service/circuitbraker"
	"limiterdemo/user_service/limiter"
	"limiterdemo/user_service/proto/user_service"
	us "limiterdemo/user_service/service"
	"log"
	"net"
	"net/http"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	opts := []sre.Option{ // 内置失败率，当滑动窗口失败总数超过总请求数比例的50%就自动触发熔断
		sre.WithRequest(100), // 设置最小请求阈值，如果请求数太少就不触发了，避免误触发
		//sre.WithBucket(10),   // 将滑动窗口时间拆分成n个桶
		//sre.WithSuccess(10),  // 半开状态的成功阈值
		//sre.WithWindow(5),    // 设置滑动窗口的总时间(统计请求数据的时间范围)
	}
	cirb := cb.NewCircuitBraker(opts...)
	userService := us.NewUserService(cirb)

	//每秒生成100个令牌，桶容量200（允许突发200个请求）
	li := limiter.NewTokenBucketLimiter(rate.Limit(100), 200)

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			li.LimitInterceptor(),
			cirb.CircuitBrakerInterceptor()),
	)

	service.RegisterUserServiceServer(s, userService)
	log.Println("Staring user service in 50051...")

	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	conn, err := grpc.NewClient("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux()
	ctx := context.Background()
	err = service.RegisterUserServiceHandler(ctx, mux, conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Staring gRPC gateway on :8081")
	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal("Failed to serve HTTP:", err)
	}
}
