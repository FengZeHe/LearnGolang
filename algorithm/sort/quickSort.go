package main

import "fmt"

func main() {
	arr := []int{4, 56, 3, 1, 5, 2}
	quickSort(arr)
	fmt.Println(arr)
}

func quickSort(arr []int) {
	arrlen := len(arr)
	if arrlen < 2 {
		return
	}
	head, pivot := 0, arrlen-1
	value := arr[head]
	for head < pivot {
		// 比pivot大的放右边
		if arr[head+1] > value {
			arr[head+1], arr[pivot] = arr[pivot], arr[head+1]
			pivot--
		} else if arr[head+1] < arr[head] {
			//如果元素遇到比它小的 就更换位置
			arr[head], arr[head+1] = arr[head+1], arr[head]
			head++
		} else {
			head++
		}
		quickSort(arr[:head])
		quickSort(arr[head+1:])
	}

}
