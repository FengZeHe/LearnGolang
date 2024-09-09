package main

/*
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

示例 1：
输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

示例 2：
输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
*/

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	bestTimeBuyStocks(prices)

}

/*
思路：
1.题目的意图是找出两数相减的最大值，而且有顺序关系，小值元素位置要在大值元素之前。
笨方法，枚举法，从后往前相减，找出值最小的数
*/
func bestTimeBuyStocks(prices []int) (ans int) {
	minPrice := prices[0]
	for _, p := range prices {
		ans = max(ans, p-minPrice)
		minPrice = min(minPrice, p)
	}
	return ans
}
