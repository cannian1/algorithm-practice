# 239. 滑动窗口最大值
# https://leetcode.cn/problems/sliding-window-maximum/

from typing import List

from collections import deque


class MyQueue:  # 单调队列（从大到小）
    """假定右边是入口，左边是出口"""

    def __init__(self):
        self.queue = deque()

    # 每次弹出时，比较当前要弹出的数值是否等于队列出口元素的数值，如果相等则弹出。
    # 同时 pop 之前判断队列当前是否为空。
    def pop(self, value):
        if self.queue and value == self.queue[0]:
            self.queue.popleft()

    # 如果 push 的数大于入口元素的数值，那么就将队列左边的数值弹出，直到 push 的数值小于等于队列入口处元素的数值为止
    # 这样就保持了队列里的数值是单调从大到小的了
    def push(self, value):
        while self.queue and value > self.queue[-1]:
            self.queue.pop()
        self.queue.append(value)

    # 查询当前队列里的最大值，直接返回队列前端的值就可以
    def front(self):
        return self.queue[0]


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        que = MyQueue()
        result = []
        for i in range(k):  # 先将前 k 个元素放进队列
            que.push(nums[i])
        result.append(que.front())  # result 记录前 k 个元素的最大值
        for i in range(k, len(nums)):
            que.pop(nums[i - k])  # 滑动窗口移除最前面元素
            que.push(nums[i])  # 滑动窗口加入最后面的元素
            result.append(que.front())  # 记录对应最大值
        return result
