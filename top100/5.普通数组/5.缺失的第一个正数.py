# 41. 缺失的第一个正数 https://leetcode.cn/problems/first-missing-positive/
from typing import List

# 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

"""
如果本题没有额外的时空复杂度要求，那么就很容易实现：

我们可以将数组所有的数放入哈希表，随后从 1 开始依次枚举正整数，并判断其是否在哈希表中； O(n),O(n)
我们可以从 1 开始依次枚举正整数，并遍历数组，判断其是否在数组中。 O(n²),O(1)
"""

"""
实际上，对于一个长度为 N 的数组，其中没有出现的最小正整数只能在 [1,N+1] 中。这是因为如果 [1,N] 都出现了，那么答案是 N+1，
否则答案是 [1,N] 中没有出现的最小正整数。
"""


class Solution:
    """哈希表"""
    """
    最后的结果一定是在 [1,n+1] 内
    对应的下标 [0,n] 里 nums[i] 哪个不是负数（没有做标记），i+1就是答案
    """

    def firstMissingPositive(self, nums: List[int]) -> int:
        n = len(nums)
        # 将所有负数和零过滤（不可能为结果）变为一个足够大的正数
        for i in range(n):
            if nums[i] <= 0:
                nums[i] = n + 1

        # 重排数组：再次遍历数组，对于每个元素nums[i]，如果它的绝对值在1到n之间，那么我们将索引为nums[i] - 1的元素变为负数，
        # 表示数字nums[i]已经在数组中找到了正确的位置。
        for i in range(n):
            num = abs(nums[i])  # 这一步使用的是绝对值，因为之前的元素可能已经被标记为负数
            if num <= n:
                nums[num - 1] = -abs(nums[num - 1])

        for i in range(n):
            if nums[i] > 0:
                return i + 1

        return n + 1


arr = [3, 4, -1, 1]
s = Solution()
res = s.firstMissingPositive(arr)
print(res)
