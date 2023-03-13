package main

import "fmt"

func selectionSort(nums []int) {
	if (len(nums)) <= 1 {
		return
	}

	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}

		}
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]

		}

	}
	fmt.Println(nums)
}

func main() {
	nums := []int{4, 5, 6, 7, 8, 3, 2, 1}
	selectionSort(nums)
}
