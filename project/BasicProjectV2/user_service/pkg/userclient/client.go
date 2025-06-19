package userclient

import (
	"context"
	"github.com/basicprojectv2/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type UserServiceClient struct {
	client user_service.UserServiceClient
	conn   *grpc.ClientConn
}

func NewUserServiceClient(addr string) (*UserServiceClient, error) {
	// 建立不安全连接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}

	return &UserServiceClient{
		client: user_service.NewUserServiceClient(conn),
		conn:   conn,
	}, nil

}

func (c *UserServiceClient) Close() error {
	return c.conn.Close()
}

func (c *UserServiceClient) GetUserById(ctx context.Context, id string) (*user_service.User, error) {
	req := &user_service.GetUserByIdReq{Id: id}
	res, err := c.client.GetUserById(ctx, req)
	if err != nil {
		log.Printf("GetUserById: %v", err)
	}
	return res, err
}
