# 347. 前 K 个高频元素
# https://leetcode.cn/problems/top-k-frequent-elements/

import heapq
from typing import List


# 时间复杂度：O(nlogk)
# 空间复杂度：O(n)
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        # 要统计元素出现频率
        map_ = {}  # nums[i]:对应出现的次数 map的 key 是元素，值是频率
        for i in range(len(nums)):
            map_[nums[i]] = map_.get(nums[i], 0) + 1

        # 对频率排序
        # 定义一个小顶堆，大小为 k
        pri_que = []  # 小顶堆

        # 用固定大小为 k 的小顶堆，扫描所有频率的数值
        for key, freq in map_.items():
            heapq.heappush(pri_que, (freq, key))  # 最小堆存储 元组(频率，元素)
            if len(pri_que) > k:  # 如果堆的大小大于了k，则队列弹出，保证堆的大小一直为k
                heapq.heappop(pri_que)

        # 找出前 k 个高频元素
        result = [0] * k
        # 遍历区间为 [k-1, -1) 遍历步长为 -1
        for i in range(k - 1, -1, -1):
            result[i] = heapq.heappop(pri_que)[1]
        return result


sol = Solution()

# 测试用例 1
nums = [1, 1, 1, 2, 2, 3]
k = 2
r = sol.topKFrequent(nums, k)
print(r)

# 测试用例 2
nums = [1]
k = 1
r = sol.topKFrequent(nums, k)
print(r)

# 测试用例 3
nums = [1, 2]
k = 2
r = sol.topKFrequent(nums, k)
print(r)

# 测试用例 4
nums = [4, 1, -1, 2, -1, 2, 3]
k = 2
r = sol.topKFrequent(nums, k)
print(r)
