package main

import "fmt"

// defer会在return和panic前执行，相当于java中finally的概念 ，执行顺序类似于栈（先defer后执行）
func main() {
	fmt.Println("hello")
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("world")
	defer fmt.Println("3")

}
