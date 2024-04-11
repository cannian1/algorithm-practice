# 11. 盛最多水的容器 https://leetcode.cn/problems/container-with-most-water
from typing import List


class Solution:
    """容器最大容积从两侧向中间遍历"""
    def maxArea(self, height: List[int]) -> int:
        left, right = 0, len(height) - 1
        result = 0
        while left < right:
            if height[left] < height[right]:
                result = max(result, height[left] * (right - left))
                left += 1
            else:
                result = max(result, height[right] * (right - left))
                right -= 1
        return result
