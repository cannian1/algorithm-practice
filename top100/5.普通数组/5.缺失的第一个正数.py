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
    """哈希表（看不明白）"""
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
        # 表示数字nums[i]已经在数组中找到了正确的位置。（[1,n+1] 范围的元素的索引都减一的位置一定都能在原列表中找到）
        for i in range(n):
            num = abs(nums[i])  # 这一步使用的是绝对值，因为之前的元素可能已经被标记为负数
            if num <= n:
                nums[num - 1] = -abs(nums[num - 1])

        for i in range(n):
            if nums[i] > 0:
                return i + 1

        return n + 1


class Solution2:
    """置换（好理解）"""
    """
    如果数组中包含 x∈[1,N]，那么恢复后，数组的第 x−1 个元素为 x。
    在恢复后，数组应当有 [1, 2, ..., N] 的形式，但其中有若干个位置上的数是错误的，每一个错误的位置就代表了一个缺失的正数。
    以题目中的示例二 [3, 4, -1, 1] 为例，恢复后的数组应当为 [1, -1, 3, 4]，我们就可以知道缺失的数为 2
    
    
    我们可以对数组进行一次遍历，对于遍历到的数 x=nums[i]，如果 x∈[1,N]，我们就知道 x 应当出现在数组中的 x−1 的位置，
    因此交换 nums[i] 和 nums[x−1]，这样 x 就出现在了正确的位置。在完成交换后，新的 nums[i] 可能还在 [1,N] 的范围内，
    我们需要继续进行交换操作，直到 x∉[1,N]。
    
    上面的方法可能会陷入死循环。如果 nums[i] 恰好与 nums[x−1] 相等，那么就会无限交换下去。
    此时我们有 nums[i]=x=nums[x−1]，说明 xxx 已经出现在了正确的位置。因此我们可以跳出循环，开始遍历下一个数。
    """

    def firstMissingPositive(self, nums: List[int]) -> int:
        n = len(nums)
        for i in range(n):
            while 1 <= nums[i] <= n and nums[nums[i] - 1] != nums[i]:
                nums[nums[i] - 1], nums[i] = nums[i], nums[nums[i] - 1]

        for i in range(n):
            if nums[i] != i + 1:
                return i + 1
        return n + 1


arr = [3, 4, -1, 1]
s = Solution()
res = s.firstMissingPositive(arr)
print(res)
