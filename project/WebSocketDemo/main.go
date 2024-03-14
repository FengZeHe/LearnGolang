package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	Cors "github.com/websocketdemo/Middleware/cors"
	"log"
	"net/http"
	"sync"
	"time"
)

var UpGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// 存储连接的映射
var connections = make(map[string]*websocket.Conn)
var mu sync.Mutex

func WebSocketHandler(c *gin.Context) {
	//clientId := ws.RemoteAddr().String()
	//token := c.GetHeader("auth")
	userID := c.Query("userid")
	log.Println("userID->", userID)
	// 获取WebSocket连接 下面这行代码使用的方法已经被弃用
	//ws, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	ws, err := UpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	// 处理WebSocket消息
	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("msg:", string(p))

		mu.Lock()
		connections[userID] = ws
		mu.Unlock()
		err = ws.WriteMessage(messageType, []byte("copy"))
		if err != nil {
			log.Println("websocket write Message ERROR")
		}
	}

	// 关闭WebSocket连接
	defer ws.Close()
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

	go NotificationService()

	r.Run(":8899") // 在0.0.0.0:8899上监听并服务
}

// 定时发送
func NotificationService() {
	for range time.Tick(5 * time.Second) {
		for clientID, conn := range connections {
			fmt.Println(clientID)
			_ = conn.WriteMessage(1, []byte("Notification Service Msg"))
		}
		fmt.Println("Send Message")
	}

}
