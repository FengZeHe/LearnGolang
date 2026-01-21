package events

import (
	"log"
	"time"
)

func ExecTimeKeeping(taskID uint) (err error) {
	log.Println("正在执行报时任务", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}
