# 121.买卖股票的最佳时机
# https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        ans = 0
        min_pirce = prices[0]  # min_pirce 维护的是 prices[i] 左侧元素的最小值
        for p in prices:
            ans = max(ans, p - min_pirce)  # 由于只能买卖一次，所以在遍历中，维护 prices[i] - min_pirce 的最大值就是答案
            min_pirce = min(min_pirce, p)
        return ans
