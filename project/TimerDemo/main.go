package main

import (
	"log"
	"time"
)

func main() {

	// 创建每1s触发一次的Ticker
	ticker := time.NewTicker(1 * time.Second)

	// 创建倒计时5s的timer
	timer := time.NewTimer(5 * time.Second)

	go func() {
		log.Println("2s 开始")
		time.Sleep(2 * time.Second)
		log.Println("2s 结束")
	}()

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("ticker 1s定时器")
			case <-timer.C:
				log.Println("5s到了")
				ticker.Stop()
				done <- true
				return

			}
		}
	}()

	<-done
	log.Println("done")

}
