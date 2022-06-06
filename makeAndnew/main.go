package main

import "fmt"

func main() {
	var num *int
	num = new(int)
	*num = 100
	fmt.Println(*num, num)
	// 100 0xc0000180a0

	myslice1 := make([]int, 0)
	myslice2 := make([]int, 0)
	myslice3 := make([]int, 10)
	myslice4 := make([]int, 10, 20)

	fmt.Println(myslice1, myslice2, myslice3, myslice4)
}
