package main

import "fmt"

func main() {
	defer func() {
		if data := recover(); data != nil {
			fmt.Println("恢复回来了")
		}
	}()

	panic("holy shit")
	fmt.Println("不会执行到这里")
}
