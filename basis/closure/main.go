package main

import (
	"fmt"
)

// 函数闭包 = 匿名函数 + 定义它的上下文
func main() {
	value := test()
	fmt.Println(value())
	fmt.Println(value())
	fmt.Println(value())
	fmt.Println(value())
	fmt.Println(value())

	fmt.Println(ReturnClosure("Tim")())

	Delay()
}

/*
闭包的用法
*/
func ReturnClosure(name string) func() string {
	return func() string {
		word := fmt.Sprintf("holy shit! You are %s !! ", name)
		return word
	}
}

/*
闭包延时绑定

	闭包保存/记录了它产生时的外部函数的所有环境。如同普通变量/函数的定义和实际赋值/调用或者说执行，是两个阶段

闭包也是一样， for循环中仅仅声明了恶一个闭包 返回的fns是一段闭包函数定义，只有在外部执行了fn（）时才真正执行了碧波啊，
在执行这个碧波啊的时候，会去其外部环境解引用，这时候四个函数引用的都是同一个变量的内存地址，因为变量值已经改变，所以得到的值是最后最新的值
*/
func Delay() {
	fmt.Println("delay begin")
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
		fns = append(fns, func() {
			fmt.Println("hello，this is", i)
		})
	}

	for _, fn := range fns {
		//这里调用了才执行
		fn()
	}

}

func test() func() int {
	var a int
	return func() int {
		a++
		return a
	}
}
