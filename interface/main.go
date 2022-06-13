package main

import (
	"fmt"
)

// 定义接口
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
