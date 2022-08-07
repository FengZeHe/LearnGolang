package main

import "fmt"

var initvalue = 0

func init() {
	initvalue = 1
	fmt.Println("init function has been run...")
}

func main() {
	fmt.Println(initvalue)
}
