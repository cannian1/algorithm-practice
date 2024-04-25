# 240. 搜索二维矩阵 II
# https://leetcode.cn/problems/search-a-2d-matrix-ii/
from typing import List


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        """以右上角为起点，用二叉搜索树的方式找"""
        m, n = len(matrix), len(matrix[0])
        x, y = 0, n - 1
        while x < m and y >= 0:
            if matrix[x][y] == target:
                return True
            if matrix[x][y] > target:
                y -= 1
            if matrix[x][y] < target:
                x += 1
        return False
