# 80.删除
# https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/
from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        slow = 0
        for fast in range(len(nums)):
            if slow < 2 or nums[fast] != nums[slow - 2]:
                nums[slow] = nums[fast]
                slow += 1
        return slow