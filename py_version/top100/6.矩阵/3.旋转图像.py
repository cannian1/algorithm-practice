# 48. 旋转图像 https://leetcode.cn/problems/rotate-image/
from typing import List


class Solution0:
    """使用辅助数组 O(n²) O(n²)"""

    def rotate(self, matrix: List[List[int]]) -> None:
        n = len(matrix)
        matrix_new = [[0] * n for _ in range(n)]
        for i in range(n):
            for j in range(n):
                matrix_new[j][n - i - 1] = matrix[i][j]
        matrix[:] = matrix_new


class Solution:
    """翻转代替旋转 O(n²) O(1)"""
    """先水平翻转，然后沿主对角线翻转"""

    def rotate(self, matrix: List[List[int]]) -> None:
        n = len(matrix)
        # 水平翻转
        for i in range(n // 2):
            for j in range(n):
                matrix[i][j], matrix[n - i - 1][j] = matrix[n - i - 1][j], matrix[i][j]

        # 主对角线翻转
        for i in range(n):
            for j in range(i):
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]


s = Solution()
arr = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
s.rotate(arr)
print(arr)
