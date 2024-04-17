# 1. 两数之和
# https://leetcode.cn/problems/two-sum/

from typing import List

"""
[2,7,3,6] target = 9
假设当前遍历到了 3，那么需要检查前面是否遍历过 6，因为 3+6 = 9

因为要 ‘查找’ 的是元素，所以要把元素作为 map 的 key，索引作为 map 的 value

"""


class Solution:
    """使用 dict"""

    def twoSum(self, nums: List[int], target: int) -> List[int]:
        record = dict()
        for i, v in enumerate(nums):
            if target - v in record:
                return [record[target - v], i]
            record[v] = i  # 把遍历过的元素加到 map 里，map 的 key是值，map的 value 是索引
        return []


class Solution2:
    """使用 set"""

    def twoSum(self, nums: List[int], target: int) -> List[int]:
        seen = set()
        for i, num in enumerate(nums):
            complement = target - num
            if complement in seen:
                return [nums.index(complement), i]
            seen.add(num)
