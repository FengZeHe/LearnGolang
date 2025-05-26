package main

import (
	"github.com/golang/protobuf/proto"
	"log"
	"protobufdemo/service"
)

// protoc --go_out=./ .user.proto
func main() {
	user1 := &service.User{
		Userid:   123456,
		Username: "user001",
		Age:      18,
	}
	marshal, err := proto.Marshal(user1)
	if err != nil {
		panic(err)
	}

	user2 := &service.User{}
	if unErr := proto.Unmarshal(marshal, user2); unErr != nil {
		panic(unErr)
	}
	log.Println(user2.String())

}
