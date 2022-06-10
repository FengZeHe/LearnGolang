package main

import (
	"fmt"

	_ "github.com/LearnGolang/funcreturn/moudleA"
	_ "github.com/LearnGolang/funcreturn/moudleB"
)

func main() {
	fmt.Println("hello from main")
}

func init() {
	fmt.Println("init from main")
}
