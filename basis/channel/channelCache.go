package main

import (
	"fmt"
	"time"
)

/*
1. 不带缓冲的channel：发送和接接收动作是同时发生的，如果没有读取就会一直阻塞；就像到付的快递，快递员在门口等你等你付了邮费快递员才走。
2. 有缓冲是异步的，类似于一个队列，只有队排完了才可能发生阻塞，就像快递员将快递放在你家门口，送完就走。除非你家门口满了塞不下，快递员就会等你家门口空下来再送快递到你家。
*/
func main() {
	channelWithCache()
	channelWithoutCache()
}

// 有缓冲的
func channelWithCache() {
	ch := make(chan string, 1)
	go func() {
		ch <- "holy shit"
		time.Sleep(time.Second)
		ch <- "hello"
	}()

	time.Sleep(time.Second * 2)
	msg := <-ch
	fmt.Println(time.Now().String(), msg)
	msg = <-ch
	fmt.Println(time.Now().String(), msg)
}

// 不带缓冲
func channelWithoutCache() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second)
		ch <- "run run run "
	}()
	msg := <-ch
	fmt.Println("msg:", msg)
}
