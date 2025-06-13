package main

import (
	"context"
	"github.com/basicprojectv2/pkg/userclient"
	"log"
)

func main() {
	c, err := userclient.NewUserServiceClient("localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// 调用
	user, err := c.GetUserById(context.Background(), "1821841651400708096")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User --> ", user)
}
