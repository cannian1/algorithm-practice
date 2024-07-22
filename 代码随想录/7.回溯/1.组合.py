# 77.组合
# https://leetcode.cn/problems/combinations/
from typing import List


class Solution:
    """未剪枝优化"""

    def combine(self, n: int, k: int) -> List[List[int]]:
        result = []  # 存放结果集
        self.backtracking(n, k, 1, [], result)
        return result

    def backtracking(self, n, k, startIndex, path, result):
        if len(path) == k:
            result.append(path[:])
            return
        for i in range(startIndex, n + 1):  # 需要优化的地方
            path.append(i)  # 处理节点
            self.backtracking(n, k, i + 1, path, result)
            path.pop()  # 回溯，撤销处理的节点


class Solution2:
    """剪枝优化"""

    def combine(self, n: int, k: int) -> List[List[int]]:
        result = []  # 存放结果集
        self.backtracking(n, k, 1, [], result)
        return result

    def backtracking(self, n, k, startIndex, path, result):
        """
        startIndex: 搜索的起始位置
        """
        # 到树形结构的叶子节点了，收集结果
        if len(path) == k:
            result.append(path[:])
            return
        # 如果for循环选择的起始位置之后的元素个数 已经不足 我们需要的元素个数了，那么就没有必要搜索了
        # 假设 n = 4, k = 3。在 path的大小为 0 的时候，4 - 3 + 1 = 2。即从 2 开始搜索还是合理的，可以组合出 [2,3,4]
        for i in range(startIndex, n - (k - len(path)) + 2):  # 需要优化的地方 i <= n - (k - len(path)) + 1
            path.append(i)  # 处理节点
            self.backtracking(n, k, i + 1, path, result)
            path.pop()  # 回溯，撤销处理的节点
