// 121.买卖股票的最佳时机
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

package greedy

func maxProfit(prices []int) int {
	res := 0
	minPrice := prices[0] // minPrice 维护的是 price[i] 左侧元素的最小值
	for _, p := range prices {
		// 由于只能买卖一次，所以在遍历中，维护 prices[i] - minPrice 的最大值就是答案
		res = max(res, p-minPrice)
		minPrice = min(minPrice, p)
	}
	return res
}
