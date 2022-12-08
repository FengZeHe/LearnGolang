package main

import (
	"errors"
	"fmt"
)

/*
error 一般用于表达可以被处理的错误，而panic用于表达不可恢复的错误
*/
func main() {
	err := errors.New("new error")
	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2: [%w]", err1)
	err3 := fmt.Errorf("err3: [%w]", err2)
	fmt.Println("1:", err1, "2:", err2, "3:", err3)

	res := errors.Is(err1, err)
	fmt.Println(res)
	if err2 == errors.Unwrap(err3) {
		fmt.Println("unwrapped !!!")
	}

	//如果err2/err3是被包过的话，那么err1是用来找出来最里面的err，还是一层一层的找
	if errors.Is(err3, err1) {
		fmt.Println("wrapped is err")
	}

	/*
		As找到err链中与target匹配的第一个错误，如果找到则将target设置为错误值并返回true,否则返回false
	*/
	fmt.Println(errors.As(err3, &err))
}
