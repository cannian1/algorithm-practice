# 797.所有可能的路径
# https://leetcode.cn/problems/all-paths-from-source-to-target/
from typing import List


class Solution:
    def __init__(self):
        self.result = []
        self.path = [0]

    def allPathsSourceTarget(self, graph: List[List[int]]) -> List[List[int]]:
        if not graph:
            return []

        self.dfs(graph, 0)
        return self.result

    def dfs(self, graph, root: int):
        if root == len(graph) - 1:  # 成功找到一条路径
            self.result.append(self.path[:])
            return

        for node in graph[root]:  # 遍历节点 n 的所有后序节点
            self.path.append(node)
            self.dfs(graph, node)
            self.path.pop()  # 回溯
