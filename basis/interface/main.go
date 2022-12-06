package main

import (
	"fmt"
)

// 定义接口
/*
	Interface 里面只能有方法，方法也不需要func 关键字
	接口是一组行为的抽象， 在编程中尽量使用接口  面向接口编程
*/
type Phone interface {
	call()
}

// 定义结构体
type NokiaPhone struct {
}

// 实现接口方法
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia ,i can call you ")
}

type IPhone struct {
}

func (iphone IPhone) call() {
	fmt.Println("I am iphone,i can call you either")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
