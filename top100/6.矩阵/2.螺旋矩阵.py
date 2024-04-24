# 54. 螺旋矩阵 https://leetcode.cn/problems/spiral-matrix/
from typing import List


class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        if not matrix or not matrix[0]:
            return []

        left, right, top, bottom, res = 0, len(matrix[0]) - 1, 0, len(matrix) - 1, []
        while True:
            for i in range(left, right + 1):
                res.append(matrix[top][i])  # left to right
            top += 1
            if top > bottom:
                break
            for i in range(top, bottom + 1):
                res.append(matrix[i][right])  # top to bottom
            right -= 1
            if left > right:
                break
            for i in range(right, left - 1, -1):
                res.append(matrix[bottom][i])  # right to left
            bottom -= 1
            if top > bottom:
                break
            for i in range(bottom, top - 1, -1):
                res.append(matrix[i][left])  # bottom to top
            left += 1
            if left > right:
                break
        return res
