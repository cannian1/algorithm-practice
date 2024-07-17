# 39.组合总合
# https://leetcode.cn/problems/combination-sum/
from typing import List


class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        result = []
        self.backtracking(candidates, target, 0, 0, [], result)
        return result

    def backtracking(self, candidates, target, total, startIndex, path, result):
        if total > target:
            return
        if total == target:
            result.append(path[:])
            return

        for i in range(startIndex, len(candidates)):
            total += candidates[i]
            path.append(candidates[i])
            self.backtracking(candidates, target, total, i, path, result)
            total -= candidates[i]
            path.pop()


class BetterSolution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        result = []
        candidates.sort()
        self.backtracking(candidates, target, 0, 0, [], result)
        return result

    def backtracking(self, candidates, target, total, startIndex, path, result):
        if total == target:
            result.append(path[:])
            return

        for i in range(startIndex, len(candidates)):
            if total + candidates[i] > target:
                continue
            total += candidates[i]
            path.append(candidates[i])
            self.backtracking(candidates, target, total, i, path, result)
            total -= candidates[i]
            path.pop()


# 补充：77.组合
class SolutionCombine:
    def combine(self, n: int, k: int) -> List[List[int]]:
        result = []
        path = []  # [0]*k

        def dfs(i: int) -> None:
            # 剪枝
            d = k - len(path)
            if i < d:
                return

            if len(path) == k:
                result.append(path.copy())
                return

            for j in range(i, 0, -1):
                path.append(j)
                dfs(j - 1)
                path.pop()
        dfs(n)
        return result