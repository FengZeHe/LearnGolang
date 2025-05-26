package main

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"protobufdemo/service"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("连接服务端失败: %v", err)
	}
	defer conn.Close()
	log.Println("已连接到服务端")

	user := &service.User{
		Userid:   123456,
		Username: "长不大的志明",
		Age:      29,
	}

	msgData, err := proto.Marshal(user)
	if err != nil {
		log.Fatalf("序列化失败: %v", err)
	}

	lenBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBuf, uint32(len(msgData)))

	if _, err := conn.Write(lenBuf); err != nil {
		log.Fatalf("发送长度失败: %v", err)
	}
	if _, err := conn.Write(msgData); err != nil {
		log.Fatalf("发送消息失败: %v", err)
	}

	log.Println("User 消息发送成功")
}
