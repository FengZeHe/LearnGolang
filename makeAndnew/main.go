package main

import "fmt"

func main() {
	var num *int
	num = new(int)
	*num = 100
	fmt.Println(*num, num)
	// 100 0xc0000180a0
}
