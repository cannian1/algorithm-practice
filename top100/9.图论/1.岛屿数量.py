# 200.岛屿数量
# https://leetcode.cn/problems/number-of-islands/
from typing import List


# 我们用三个状态去标记每一个格子
# 0 代表海水
# 1 代表陆地
# 2 代表已经访问的陆地
class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        res = 0

        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] == "1":
                    res += 1
                    self.traversal(grid, i, j)

        return res

    def traversal(self, grid, i, j):
        m = len(grid)
        n = len(grid[0])

        if i < 0 or j < 0 or i >= m or j >= n:
            return  # 越界了
        elif grid[i][j] == "2" or grid[i][j] == "0":
            return

        grid[i][j] = "2"
        self.traversal(grid, i - 1, j)  # 往上走
        self.traversal(grid, i + 1, j)  # 往下走
        self.traversal(grid, i, j - 1)  # 往左走
        self.traversal(grid, i, j + 1)  # 往右走
