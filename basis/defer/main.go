package main

import "fmt"

func main() {
	fmt.Println("hello")
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("world")
	defer fmt.Println("3")

}
