package main

import "fmt"

func main() {
	arr := []int{4, 5, 1, 2, 3, 6}
	InsertSort(arr)
}

func InsertSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 0; i < len(arr); i++ {
		value := arr[i]
		j := i - 1
		//查找插入位置
		for ; j >= 0; j-- {
			if arr[j] > value {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		// 插入数据
		arr[j+1] = value
	}
	fmt.Println(arr)

}
