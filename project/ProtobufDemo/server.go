package main

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"protobufdemo/service"
)

func main() {
	listener, listenerErr := net.Listen("tcp", ":8080")
	if listenerErr != nil {
		log.Fatalf("监听端口失败", listenerErr)
	}
	defer listener.Close()
	log.Println("服务启动...")

	for {
		conn, connErr := listener.Accept()
		if connErr != nil {
			log.Fatalf("connect error", connErr)
		}
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("connection close error", err)
		}
	}(conn)
	log.Println("已连接客户端", conn.RemoteAddr())

	// 步骤 1：读取消息长度（前 4 字节表示长度，大端序）
	lenBuf := make([]byte, 4)
	if _, err := conn.Read(lenBuf); err != nil {
		log.Printf("读取长度失败: %v", err)
		return
	}
	msgLen := binary.BigEndian.Uint32(lenBuf) // 转换为 uint32 类型的长度

	msgBuf := make([]byte, msgLen)
	if _, err := conn.Read(msgBuf); err != nil {
		log.Printf("读取消息失败: %v", err)
		return
	}

	// 步骤 3：反序列化 Protobuf 消息
	user := &service.User{}
	if err := proto.Unmarshal(msgBuf, user); err != nil {
		log.Printf("反序列化失败: %v", err)
		return
	}

	// 输出解析结果
	log.Printf("收到 User 消息：\nuserid: %d\nusername: %s\nage: %d",
		user.Userid, user.Username, user.Age)
}
