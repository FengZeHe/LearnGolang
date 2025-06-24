package main

import (
	"context"
	pb "github.com/basicprojectv2/proto/user_service"
	"github.com/basicprojectv2/user_client/interceptors/jwt"
	lg "github.com/basicprojectv2/user_client/interceptors/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func main() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3ODY5NjU0MDAsIlVzZXJJZCI6IjE5MzQ4ODI1NjQ3MDM1MjI4MTYifQ.UvpMjVj5AsYemY0X9HXvDr-BOlLGgniklZb45_qg8_g"
	ji := jwt.NewJWTInterceptor(token)

	li := lg.NewLogInterceptor()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			ji.JwtClientInterceptor(),
			li.LogUnaryServerInterceptor()))
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
