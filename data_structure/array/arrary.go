package main

import "fmt"

func main() {
	var arr = [10]int{}
	for i := 0; i < len(arr); i++ {
		fmt.Print(&arr[i], " ")
	}

}
