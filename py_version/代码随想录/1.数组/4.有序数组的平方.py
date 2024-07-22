# leetcode 997.有序数组的平方
# https://leetcode.cn/problems/squares-of-a-sorted-array/
from typing import List

# 有序数组平方后的最大值在数组的两端 （最左侧的负数平方后可能变成最大的元素）
# 双指针

class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        k = len(nums) - 1
        new_nums = [float('inf')] * len(nums)
        left, right = 0, len(nums) - 1
        while left <= right:
            # 比较两侧元素的平方大小
            if nums[left] ** 2 < nums[right] ** 2:
                # 新的数组的元素更新为 两侧元素大的那方
                new_nums[k] = nums[right] ** 2
                right -= 1
            else:
                new_nums[k] = nums[left] ** 2
                left += 1
            k -= 1
        return new_nums

    def sortedSquaresBad(self, nums: List[int]) -> List[int]:
        new_nums = [0] * len(nums)
        for i, v in enumerate(nums):
            new_nums[i] = v * v
        new_nums.sort()
        return new_nums

