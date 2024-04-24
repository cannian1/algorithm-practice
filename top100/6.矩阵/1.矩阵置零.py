# 73. 矩阵置零 https://leetcode.cn/problems/set-matrix-zeroes/
from typing import List


class Solution:
    """使用标记数组 O(mn) O(m+n)"""

    def setZeroes(self, matrix: List[List[int]]) -> None:
        m, n = len(matrix), len(matrix[0])
        row, col = [False] * m, [False] * n

        # 先遍历这个矩阵，遇到 0 就将该元素所在的行和列对应标记数组的位置置为 true
        for i in range(m):
            for j in range(n):
                if matrix[i][j] == 0:
                    row[i] = col[j] = True

        for i in range(m):
            for j in range(n):
                if row[i] or col[j]:
                    matrix[i][j] = 0


class Solution2:
    """只用一个标记变量标记行首，其他用行首和列首是否为0来判断是否更改对应行和列 O(mn) O(1)"""

    def setZeroes(self, matrix: List[List[int]]) -> None:
        m, n = len(matrix), len(matrix[0])
        first_row = False  # 标记首行是否有 0 元素

        for i, row in enumerate(matrix):
            for j, item in enumerate(row):
                if i == 0 and item == 0:
                    first_row = True  # 首行出现 0元素，用标志位标记
                elif item == 0:
                    matrix[i][0] = 0  # 非首行出现 0元素，将对应的列首置为0，说明该列要置为0
                    matrix[0][j] = 0  # 将对应的行首置为0，说明该行要置为0

        for i in range(m - 1, -1, -1):
            for j in range(n - 1, -1, -1):
                # 从最后一个元素反向遍历，避免行首和列首的信息被篡改
                if i == 0 and first_row:
                    matrix[i][j] = 0  # 首行是否置为0要看标志位
                elif i != 0 and (matrix[i][0] == 0 or matrix[0][j] == 0):
                    matrix[i][j] = 0  # 非首行元素是否置为0看行首和列首是否为0
