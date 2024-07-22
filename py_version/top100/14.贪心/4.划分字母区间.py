# 763.划分字母区间
# https://leetcode.cn/problems/partition-labels/
from typing import List


class Solution:
    def partitionLabels(self, s: str) -> List[int]:
        last = [0] * 26
        for i, ch in enumerate(s):
            last[ord(ch) - ord("a")] = i

        partition = list()
        start = end = 0
        for i, ch in enumerate(s):
            end = max(end, last[ord(ch) - ord("a")])  # 找到当前字符出现的最远位置
            if i == end:  # 如果当前位置是最远位置，表示可以分割出一个区间
                partition.append(end - start + 1)
                start = end + 1

        return partition
