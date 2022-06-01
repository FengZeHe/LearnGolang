package main

import "fmt"

func main() {
	ifelsefunc(11)
}

func ifelsefunc(num int) {
	if num == 0 {
		fmt.Println(" num = 0")
	} else if num > 0 && num < 10 {
		fmt.Println("0 < num < 10")
	} else {
		fmt.Println("...")
	}
}
