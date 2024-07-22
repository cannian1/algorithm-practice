# 136.只出现一次的数字
# https://leetcode.cn/problems/single-number/
from typing import List


# 任何数和 0 做异或运算，结果仍然是原来的数
# 任何数和其自身做异或运算，结果是 0
# 异或运算满足交换律和结合律

class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        single = 0
        for num in nums:
            single ^= num
        return single
