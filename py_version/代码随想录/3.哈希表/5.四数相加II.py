# 454. 四数相加 II
# https://leetcode.cn/problems/4sum-ii/
from typing import List

"""
A 口口口口口
B 口口口口口
C 口口口口口
D 口口口口口
这道题有四个数组，可以先遍历两个数组，把结果存在 哈希表 里
然后再去遍历剩下的两个数组，从哈希表中取出能够与已遍历的元素相加等于 0 的元素

这道题不仅要保存结果是否出现过，还要保存出现过的次数(不同索引的元素相加可能值相同，即出现多次)，
所以要选用的哈希结构是 map
"""


class Solution:
    """使用字典"""

    def fourSumCount(self, nums1: List[int], nums2: List[int], nums3: List[int], nums4: List[int]) -> int:
        count = 0
        hashmap = dict()
        for val1 in nums1:
            for val2 in nums2:
                hashmap[val1 + val2] = hashmap.get(val1 + val2, 0) + 1

        for val3 in nums3:
            for val4 in nums4:
                key = -(val3 + val4)
                if key in hashmap:
                    count += hashmap[key]  # 计数要加上前两个列表里出现过的次数

        return count


class Solution2:
    """使用 defaultdict"""
    from collections import defaultdict
    def fourSumCount(self, nums1: List[int], nums2: List[int], nums3: List[int], nums4: List[int]) -> int:
        from collections import defaultdict
        rec, count = defaultdict(lambda: 0), 0
        for val1 in nums1:
            for val2 in nums2:
                defaultdict[val1 + val2] += 1

        for val3 in nums3:
            for val4 in nums4:
                count += rec.get(-(val3 + val4), 0)

        return count
