package main

import "fmt"

func main() {
	arr := []int{11, 8, 3, 9, 7, 1, 2, 5}
	res := mergeSort(arr)
	fmt.Println(res)
}

func mergeSort(s []int) []int {
	len := len(s)
	if len == 1 {
		return s
	}

	mid := len / 2
	leftArr := mergeSort(s[:mid])
	rightArr := mergeSort(s[mid:])
	return merge(leftArr, rightArr)
}

func merge(leftArr, rightArr []int) []int {
	lLen := len(leftArr)
	rLen := len(rightArr)
	res := make([]int, 0)

	lIndex, rIndex := 0, 0
	for lIndex < lLen && rIndex < rLen {
		if leftArr[lIndex] > rightArr[rIndex] {
			res = append(res, rightArr[rIndex])
			rIndex++
		} else {
			res = append(res, leftArr[lIndex])
			lIndex++
		}
	}

	if lIndex < lLen {
		res = append(res, leftArr[lIndex:]...)
	}

	if rIndex < rLen {
		res = append(res, rightArr[rIndex:]...)
	}

	return res

}
