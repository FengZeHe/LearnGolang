package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	Cors "github.com/websocketdemo/Middleware/cors"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocketHandler(c *gin.Context) {
	// 获取WebSocket连接
	ws, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		panic(err)
	}

	// 处理WebSocket消息
	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			break
		}

		fmt.Println("messageType:", messageType)
		fmt.Println("p:", string(p))

		// 输出WebSocket消息内容
		c.Writer.Write(p)
	}

	// 关闭WebSocket连接
	ws.Close()
}

func main() {
	r := gin.Default()
	r.Use(Cors.Cors())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ws", WebSocketHandler)
	r.Run(":8899") // 在0.0.0.0:8899上监听并服务
}
