package main

import "fmt"

func main() {
	testcallback(true, decide)
}

func testcallback(value bool, f func(v bool)) {
	f(value)
}

func decide(value bool) {
	if value == true {
		fmt.Println("value is true")
	} else {
		fmt.Println("value is false")
	}
}
