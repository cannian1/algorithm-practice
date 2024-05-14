# 34.在排序数组中查找元素的第一个和最后一个位置
# https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
from typing import List


class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        start = self.binary_search(nums, target)
        if start == len(nums) or nums[start] != target:
            return [-1, -1]  # nums 中不存在 target
        # 如果 start 存在，那么 end 必定存在
        # 寻找第一个大于等于 target 的下标，减一即为最后的位置
        end = self.binary_search(nums, target + 1) - 1
        return [start, end]

    def binary_search(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = left + ((right - left) >> 1)
            if nums[mid] < target:
                left = mid + 1
            else:
                right = mid - 1
        return left
