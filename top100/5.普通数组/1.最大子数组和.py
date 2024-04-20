# 53.
from typing import List


class Solution:
    """动态规划"""

    def maxSubArray(self, nums: List[int]) -> int:
        max_ans = nums[0]
        for i in range(1, len(nums)):
            # 在原数组上暂存局部最大子数组的结果
            if nums[i] + nums[i - 1] > nums[i]:
                nums[i] += nums[i - 1]
            # 与目前为止遇到的全局最优解比较
            max_ans = max(max_ans, nums[i])
        return max_ans
