package main

import "fmt"

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
