# 128. 最长连续序列
from typing import List


class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        result = 0  # 记录最长连续序列长度
        num_set = set(nums)  # nums 中所有元素放入集合

        for num in num_set:
            # 这个数字没有前驱数才进入下面的逻辑
            if (num - 1) not in num_set:
                seq_len = 1  # 初始长度为 1
                # 统计连续序列长度
                while (num + 1) in num_set:
                    seq_len += 1
                    num += 1
                result = max(result, seq_len)
        return result
