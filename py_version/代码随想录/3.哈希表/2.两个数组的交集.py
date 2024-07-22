# 349. 两个数组的交集
# https://leetcode.cn/problems/intersection-of-two-arrays/

from typing import List


class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        from collections import defaultdict
        # 用个哈希表保存数组并储存字符出现次数
        num_dict = defaultdict(int)
        for val in nums1:
            num_dict[val] += 1

        res = []
        # 遍历第二个字符列表，判断每个字符是否在哈希表中出现过
        for val in nums2:
            # 若出现过则向交集列表添加，并删除哈希表的 key以免被重复统计
            if val in num_dict:
                res.append(val)
                del num_dict[val]

        return res


class Solution2:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        """字典+集合"""
        table = {}
        for num in nums1:
            table[num] = table.get(num, 0) + 1

        # 使用集合存储结果
        res = set()
        for num in nums2:
            if num in table:
                res.add(num)
                del table[num]

        return list(res)


class Solution3:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        """使用集合"""
        # s1 = set(nums1)
        # s2 = set(nums2)
        # return list(s1.intersection(s2))
        return list(set(nums1) & set(nums2))
