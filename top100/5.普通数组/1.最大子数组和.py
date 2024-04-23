# 53.最大子数组和 https://leetcode.cn/problems/maximum-subarray/
from typing import List


class Solution:
    """动态规划"""

    def maxSubArray(self, nums: List[int]) -> int:
        max_ans = nums[0]
        for i in range(1, len(nums)):
            # 在原数组上暂存局部最大子数组的结果，前一个元素与当前元素相加只要比当前元素大就覆盖当前元素（负数就直接被覆盖掉了）
            if nums[i] + nums[i - 1] > nums[i]:
                nums[i] += nums[i - 1]
            # 与目前为止遇到的全局最优解比较()
            max_ans = max(max_ans, nums[i])
        return max_ans


s = Solution()
arr = [5, 4, -1, -2, 9, 8]  # 遍历中变为 [5, 9, 8, 6, 15, 23]
res = s.maxSubArray(arr)
print(res)
arr = [-2, 1, -3, 4, -1, 2, 1, -5, 4]  # 遍历中变为 [-2, 1, -2, 4, 3, 5, 6, 1, 5]
res = s.maxSubArray(arr)
print(res)
