# 189. 轮转数组 https://leetcode.cn/problems/rotate-array/
from typing import List


class Solution:
    """三次反转"""

    def rotate(self, nums: List[int], k: int) -> None:
        k %= len(nums)
        self.reverse(nums, 0, len(nums) - 1)
        self.reverse(nums, 0, k - 1)
        self.reverse(nums, k, len(nums) - 1)

    def reverse(self, nums, left, right):
        while left < right:
            nums[left], nums[right] = nums[right], nums[left]
            left += 1
            right -= 1


class Solution2:
    """nums[:]赋值才是在原地操作"""
    def rotate(self, nums: List[int], k: int) -> None:
        k %= len(nums)
        nums[:] = nums[-k:] + nums[:-k]


arr = [1, 2, 3, 4, 5, 6, 7]
s = Solution()
s.rotate(arr, 3)
