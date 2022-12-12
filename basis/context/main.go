package main

import (
	"context"
	"fmt"
	"time"
)

/*
context 是 goroutine 的上下文，包含 goroutine 的运行状态、环境、现场等信息。
context 主要用来在 goroutine 之间传递上下文信息，包括：取消信号、超时时间、截止时间、k-v 等。
*/
func main() {
	//WithTimeout()
	//WithCancel()
	//WithDeadline()
	WithValue()
}

func WithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	start := time.Now().Unix()
	<-ctx.Done()
	end := time.Now().Unix()
	//输出2说明这里阻塞的两秒
	fmt.Println(end - start)
}

func WithCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-ctx.Done()
		fmt.Println("context was canceled")
	}()

	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)
}

func WithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	start := time.Now().Unix()
	<-ctx.Done()
	end := time.Now().Unix()
	fmt.Println(end - start)
}

func WithValue() {
	parentKey := "parent"
	parent := context.WithValue(context.Background(), parentKey, "this is parent")

	sonKey := "son"
	son := context.WithValue(context.Background(), sonKey, "this is son")

	if parent.Value(parentKey) == nil {
		fmt.Println("parent can not get son's key-value pair")
	}

	if val := son.Value(parentKey); val != nil {
		fmt.Println("parent can not get son's  key-value pair")
	}
	fmt.Println(parent.Value(parentKey), son.Value(sonKey))

}
