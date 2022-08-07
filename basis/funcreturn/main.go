package main

import (
	"fmt"

	_ "github.com/LearnGolang/funcreturn/moudleA"
	_ "github.com/LearnGolang/funcreturn/moudleB"
)

func main() {
	fmt.Println("hello from main")
	simgleStr := demoSingle()
	fmt.Println(simgleStr)
	multStr, multNum, multBool := demoMultiple()
	fmt.Println(multStr, multNum, multBool)
}

func init() {
	fmt.Println("init from main")
}

func demoSingle() int {
	return 1
}

func demoMultiple() (value string, num int, trueOrfalse bool) {
	return "my string", 88, true
}
