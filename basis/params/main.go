package main

import (
	"fmt"
	"reflect"
)

func main() {
	testParams(1, "one", 2, 2, 2, 2, 2, 3)
}

// Go的可传入不指定长度的参数，但要参数的类型。
func testParams(num int, str string, nums ...int) {
	fmt.Println(num, str)
	for _, v := range nums {
		fmt.Println("v =", v)
	}
	fmt.Println(nums, reflect.TypeOf(nums))
}
