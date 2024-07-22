# 216. 组合总和 III
# https://leetcode.cn/problems/combination-sum-iii/
from typing import List


class Solution:
    def combinationSum3(self, k: int, n: int) -> List[List[int]]:
        result = []
        self.backtracking(n, k, 0, 1, [], result)
        return result

    def backtracking(self, targetSum, k, currentSum, startIndex, path, result):
        if currentSum > targetSum:  # 剪枝操作
            return  # 如果path的长度等于k但currentSum不等于targetSum，则直接返回
        if len(path) == k:
            if currentSum == targetSum:
                result.append(path[:])
            return
        for i in range(startIndex, 9 - (k - len(path)) + 2):  # 剪枝
            currentSum += i
            path.append(i)
            self.backtracking(targetSum, k, currentSum, i + 1, path, result)
            currentSum -= i
            path.pop()
