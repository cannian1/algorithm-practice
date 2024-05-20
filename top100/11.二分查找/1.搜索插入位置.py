# 35.搜索插入位置
# https://leetcode.cn/problems/search-insert-position/
from typing import List


class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:

        left, right = 0, len(nums) - 1

        while left <= right:
            # python 位运算符优先级比算数运算低
            mid = left + ((right - left) >> 1)
            # mid = (left + right) // 2
            if nums[mid] < target:
                left = mid + 1
            else:
                right = mid - 1
        return left


arr = [1, 3, 5, 6]
s = Solution()
r = s.searchInsert(arr, 5)
print(r)
