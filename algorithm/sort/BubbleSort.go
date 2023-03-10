package main

import "fmt"

func BubbleSort(arr []int) {
	if len(arr) < 1 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				fmt.Println(arr[j], arr[j+1])
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
}

func main() {
	var arr = []int{4, 5, 6, 3, 2, 1}
	BubbleSort(arr)
}
