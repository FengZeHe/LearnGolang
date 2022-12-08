package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func PrintOnce() {
	once.Do(func() {
		fmt.Print("只会输出一次")
	})
}

func main() {
	PrintOnce()
	PrintOnce()
	PrintOnce()
	PrintOnce()
	
}
