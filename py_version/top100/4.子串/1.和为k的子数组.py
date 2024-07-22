# 560. 和为 K 的子数组 https://leetcode.cn/problems/subarray-sum-equals-k
from typing import List

"""前缀和数组是一个辅助数组，其中每个元素表示原数组从第一个元素到当前位置元素的和。
    例如，给定一个数组 nums = [1, 2, 3, 4]，其前缀和数组 prefix_sums 将是 [1, 3, 6, 10]。
    
    前缀和数组可以在 O(1) 时间内计算任意子数组的和。例如，如果我们想要计算 nums 中从索引 i 到 j 的子数组的和，我们可以使用前缀和数组来实现，
    只需要计算 prefix_sums[j] - prefix_sums[i-1]（如果 i 不为 0）或者 prefix_sums[j]（如果 i 为 0）。
    这样我们就不需要每次都遍历子数组来计算和，从而提高了效率。
"""


class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        count = {0: 1}  # 初始化哈希表，键为前缀和，值为出现的次数，初始化时前缀和为0的出现一次
        prefix_sum = 0  # 初始化前缀和
        result = 0  # 初始化结果

        # 假设数组的前缀和数组为prefixSum，其中prefixSum[i] 表示从数组起始位置到第i个位置的元素之和。
        # 那么对于任意的两个下标i和j（i < j），如果prefixSum[j] - prefixSum[i] = k，
        # 即从第i个位置到第j个位置的元素之和等于k，那么说明从第i + 1
        # 个位置到第j个位置的连续子数组的和为k。
        for num in nums:
            prefix_sum += num  # 更新前缀和
            if prefix_sum - k in count:  # 如果存在前缀和为 prefix_sum - k 的子数组(即存在和为 k 的子数组)
                result += count[prefix_sum - k]  # 将其出现的次数加到结果上
            count[prefix_sum] = count.get(prefix_sum, 0) + 1  # 更新哈希表
        return result


s = Solution()
c = s.subarraySum([1, 2, 3, 4], 3)
print(c)
