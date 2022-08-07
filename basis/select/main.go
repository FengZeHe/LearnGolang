package main

import (
	"fmt"
	"time"
)

func main() {
	checkChan := make(chan int, 10)
	checkChan2 := make(chan int, 10)
	var v int
	checkChan <- 1
	select {
	case v = <-checkChan:
		fmt.Println("receive checkChan, value=", v)
	case v = <-checkChan2:
		fmt.Println("receive checkChan2, value=", v)
	default:
		fmt.Println("Nothing")
	}

	// 使用Select做多个计时器
	Timer1 := time.NewTicker(time.Second * 5)
	select {
	case <-Timer1.C:
		fmt.Println("5s")
	}

}
