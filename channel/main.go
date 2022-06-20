package main

import "fmt"

func main() {
	// 初始化chan 执行类型并申请容量
	var intChan chan int
	intChan = make(chan int, 10)
	intChan <- 10
	intChan <- 20
	fmt.Println(intChan)
	fmt.Println(len(intChan))
	result := <-intChan
	fmt.Println(len(intChan))
	result2 := <-intChan
	fmt.Println(result)
	fmt.Println(len(intChan))
	fmt.Println(result2)
	fmt.Println(len(intChan))

	var strChan chan string
	strChan = make(chan string, 10)
	// 长度不能是0,不然没办法传东西进去
	strChan <- "hello"
	str := <-strChan
	fmt.Println("str = ", str)

	// 这种情况除外
	var strChan2 chan string
	strChan2 = make(chan string)
	go func() {
		strChan2 <- "world"
	}()

	fmt.Println(<-strChan2)
}
