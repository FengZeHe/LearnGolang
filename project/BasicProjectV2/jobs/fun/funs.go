package fun

import (
	"log"
	"time"
)

func ExecTimeKeeping(taskID uint) (err error) {
	log.Println("正在执行报时任务", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

func SayHi(taskID uint) error {
	log.Printf("【任务%v】Hello! 我是 SayHi 任务", taskID)
	return nil
}
