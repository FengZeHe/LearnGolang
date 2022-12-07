package main

import "fmt"

// 函数闭包 = 匿名函数 + 定义它的上下文
func main() {
	value := test()
	fmt.Println(value())
	fmt.Println(value())
	fmt.Println(value())
	fmt.Println(value())
	fmt.Println(value())

}

func test() func() int {
	var a int
	return func() int {
		a++
		return a
	}
}
