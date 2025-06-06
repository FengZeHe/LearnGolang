package main

/*
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]

提示：
2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案

在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标
*/

import (
	"fmt"
)

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
}

//思路：如果存在两数A、B之和等于给定值，那么A = 给定值 - B，这时候再循环下标看哪个值等于A
/*
fun twoSum(nums []int, target int) []int {
	res := []int{}
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if temp == nums[j] {
				res = append(res, i, j)
			}
		}
	}

	return res
}
*/
/*
	优化方法：同样先将B算出来，并使用Map存储，数值为key。在遍历时查找Map中是否存在该值B，如果存在则说明B找到了，那么将value和当前i拼接起来作为下标输出。
			否则将当前值和下标存到Map中，参与下一次遍历对比。

*/
func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	res := []int{}
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		if _, isequal := indexMap[temp]; isequal {
			res = append(res, indexMap[temp], i)
			return res
		}
		indexMap[nums[i]] = i
	}
	return nil
}
