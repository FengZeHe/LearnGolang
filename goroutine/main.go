package main

import (
	"fmt"
	"time"
)

func main() {
	test()
	time.Sleep(time.Second)
}

func test() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}
}
