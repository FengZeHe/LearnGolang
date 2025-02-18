package main

import (
	"log"
	"sort"
)

/*
169. 多数数组
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
示例 1：
输入：nums = [3,2,3]
输出：3

示例 2：
输入：nums = [2,2,1,1,1,2,2]
输出：2

提示：
n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109
*/

func main() {
	num := []int{6, 6, 6, 7, 7}
	majorityElement(num)
}

/*
	思路: 先将数组排序，有一个指针用来扫描元素1的连续长度是多少，扫描完后进行n/2计算，
*/

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)
	log.Println(nums)
	count := 1
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			count++
			res = nums[i]
		} else if nums[i] != nums[i+1] {
			// 判断上一段元素是否为多数元素
			if count > len(nums)/2 {
				log.Println("存在多数元素", res)
				return res
			}
		}
	}

	return res
}
