package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	wait group用于等待一组goroutine结束
	有这些方法：
	func (wg *WaitGroup) Add(delta int) Add 用来添加 goroutine 的个数
	func (wg *WaitGroup) Done() Done 执行一次数量减 1
	func (wg *WaitGroup) Wait() Wait 用来等待结束
*/

func main() {

	var wg sync.WaitGroup
	fmt.Printf("init:             %+v\n", wg)
	//wg.Add(10)
	for i := 1; i < 10; i++ {
		//计数加 1
		wg.Add(1)
		go func(i int) {
			fmt.Printf("goroutine%d start: %+v\n", i, wg)
			time.Sleep(10 * time.Second)
			// 计数减 1
			wg.Done()
			fmt.Printf("goroutine%d end:   %+v\n", i, wg)
		}(i)
		time.Sleep(time.Second)
	}

	// 等待执行结束
	wg.Wait()
	fmt.Printf("over:             %+v\n", wg)
}
