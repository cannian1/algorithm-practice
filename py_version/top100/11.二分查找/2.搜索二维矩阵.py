# 74.搜索二维矩阵
# https://leetcode.cn/problems/search-a-2d-matrix/
from typing import List


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        m = len(matrix)  # 行数
        n = len(matrix[0])  # 列数
        # 将矩阵展平，二分搜索[0, m*n]
        left, right = 0, m * n - 1  # 左右边界（左闭右闭）

        while left <= right:
            mid = left + ((right - left) >> 1)
            r, c = divmod(mid, n)  # 根据 mid 还原行与列的索引

            if matrix[r][c] > target:
                right = mid - 1
            elif matrix[r][c] < target:
                left = mid + 1
            else:
                return True
        return False
