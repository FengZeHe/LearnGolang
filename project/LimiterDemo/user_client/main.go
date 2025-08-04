package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "limiterdemo/user_service/proto/user_service"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	req := &pb.GetUserReq{
		Id: "666",
	}

	ctx := context.Background()
	user, err := client.GetUserById(ctx, req)
	if err != nil {
		log.Fatalf("client.GetUserById err: %v", err)
	}
	log.Println("user=>", user)
}
