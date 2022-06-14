package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer defer")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("panic!!")

}
