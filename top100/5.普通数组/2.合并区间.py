# 56. 合并区间 https://leetcode.cn/problems/merge-intervals/
from typing import List


class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        if len(intervals) == 0:
            return []

        intervals.sort(key=lambda x: x[0])  # 按照区间的左边界排序
        result = [intervals[0]]  # 第一个区间可以直接放入结果集中
        for i in range(1, len(intervals)):
            if result[-1][1] >= intervals[i][0]:  # 发现重叠区间（比较当前结果集的右边界与遍历到的元素的左边界）
                # 合并区间，只需要更新结果集最后一个区间的右边界，因为根据排序，左边界已经是最小的
                result[-1][1] = max(result[-1][1], intervals[i][1])
            else:
                result.append(intervals[i])  # 区间不重叠就追加集合

        return result

