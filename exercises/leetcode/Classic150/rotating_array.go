package main

import (
	"log"
	"slices"
)

/*
189 轮转数组
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]

*/

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
}

/*
思路: 就是将数组尾部的元素放在数组的最前面，形成一个新的数组 k就是变的次数
 1. 获取最后的元素，
 2. 将最后的元素删除
 3. 将最后的元素放在数组的开头
*/
func rotate(nums []int, k int) {
	if len(nums) > 0 {
		k = k % len(nums)
		slices.Reverse(nums)
		slices.Reverse(nums[:k])
		slices.Reverse(nums[k:])
	}
	log.Println(nums)
}
