# 最小 K 个数
# https://leetcode.cn/problems/smallest-k-lcci/

import heapq
from typing import List


class Solution:
    def smallestK(self, arr: List[int], k: int) -> List[int]:
        if k == 0:
            return list()

        # 初始化一个长度为 k 的队列，每个元素来自arr前k项元素的相反数值。构造大根堆
        hp = [-x for x in arr[:k]]
        heapq.heapify(hp)
        for i in range(k, len(arr)):
            if -hp[0] > arr[i]:
                heapq.heappop(hp)
                heapq.heappush(hp, -arr[i])
        ans = [-x for x in hp]
        return ans