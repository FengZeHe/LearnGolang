package main

import (
	"fmt"
)

func main() {
	// 有限的循环
	sum := 0
	for i := 0; i < 5; i++ {
		sum = sum + i
	}
	fmt.Println("sum = ", sum)

	// Golang没有while,使用for实现while
	num := 1
	for num < 100 {
		num += num
	}

	fmt.Println("num  = ", num)

	// 无限循环
	// for {
	// 	fmt.Println("hi")
	// }

	// for-range循环
	str := "hellogolang"
	for _, v := range str {
		fmt.Println(string(v), v)
	}
}
