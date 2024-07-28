# 121.买卖股票的最佳时机
# https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        ans = 0
        min_price = prices[0]  # min_price 维护的是 prices[i] 左侧元素的最小值
        for p in prices:
            ans = max(ans, p - min_price)  # 由于只能买卖一次，所以在遍历中，维护 prices[i] - min_price 的最大值就是答案
            min_price = min(min_price, p)
        return ans
