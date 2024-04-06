# 347. 前 K 个高频元素
# https://leetcode.cn/problems/top-k-frequent-elements/

import heapq
from typing import List


# 时间复杂度：O(nlogk)
# 空间复杂度：O(n)
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        # 要统计元素出现频率
        map_ = {}  # nums[i]:对应出现的次数
        for i in range(len(nums)):
            map_[nums[i]] = map_.get(nums[i], 0) + 1

        # 对频率排序
        # 定义一个小顶堆，大小为 k
        pri_que = []  # 小顶堆

        # 用固定大小为 k 的小顶堆，扫描所有频率的数值
        for key, freq in map_.items():
            heapq.heappush(pri_que, (freq, key))
            if len(pri_que) > k:  # 如果堆的大小大于了k，则队列弹出，保证堆的大小一直为k
                heapq.heappop(pri_que)

        # 找出前 k 个高频元素
        result = [0] * k
        for i in range(k - 1, -1, -1):
            result[i] = heapq.heappop(pri_que)[1]
        return result
