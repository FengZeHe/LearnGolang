package main

import (
	"context"
	pb "github.com/basicprojectv2/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

// 客户端拦截器
func loggingInterceptor(ctx context.Context, method string, req, reqly interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opt ...grpc.CallOption) error {
	start := time.Now()
	log.Printf("开始调用方法: %s 请求: %v", method, req)
	err := invoker(ctx, method, req, reqly, cc, opt...)
	log.Println("耗时", time.Since(start))

	return err
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(loggingInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	resp, err := client.Hi(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(resp.Msg)
}
