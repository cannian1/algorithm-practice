# 283. 移动零 https://leetcode.cn/problems/move-zeroes/
from typing import List


class Solution:
    """双指针，交换0与遍历过程中的非0元素"""

    def moveZeroes(self, nums: List[int]) -> None:
        k = 0
        for i in range(len(nums)):
            if nums[i] != 0:
                nums[i], nums[k] = nums[k], nums[i]
                k += 1


arr = [1, 2, 0, 0, 4, 0, 5]
s = Solution()
s.moveZeroes(arr)
print(arr)
