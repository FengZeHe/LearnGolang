package log

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

type LogInterceptor struct {
}

func NewLogInterceptor() *LogInterceptor {
	return &LogInterceptor{}
}

func (l *LogInterceptor) LogUnaryServerInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		log.Printf("开始调用方法: %s 请求: %v", method, req)
		err := invoker(ctx, method, req, reply, cc, opts...)
		log.Println("耗时", time.Since(start))
		return err
	}
}
