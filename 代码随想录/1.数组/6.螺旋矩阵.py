# LeetCode 59. 螺旋矩阵 II
# https://leetcode.cn/problems/spiral-matrix-ii

from typing import List


# 要抓住循环不变量，每次处理的边界条件要控制住，即每条边的处理规则要相同


class Solution:
    def generateMatrix(self, n: int) -> List[List[int]]:
        # [0] * n 创建一个列表，包含n个元素，每个元素都是0
        # 然后用列表推导式把上面的操作重复 n 次并组合成一个新的列表，形成二维列表（矩阵）
        nums = [[0] * n for _ in range(n)]
        start_x, start_y = 0, 0  # 起始点

        # 找规律，n是偶数时，转的圈数是 n/2； n是奇数时，转的圈数是 n//2 （3//2=1），但是最后会留下一个中心点没有赋值，要单独处理
        loop, mid = n // 2, n // 2  # 迭代次数、n为奇数时，矩阵的中心点
        count = 1  # 计数

        for offset in range(1, loop + 1):  # 每循环一层偏移量加1，偏移量从1开始
            for i in range(start_y, n - offset):  # 从左至右，左闭右开（不包含本行最后一个元素）
                nums[start_x][i] = count
                count += 1
            for i in range(start_x, n - offset):  # 从上至下
                nums[i][n - offset] = count
                count += 1
            for i in range(n - offset, start_y, -1):  # 从右至左
                nums[n - offset][i] = count
                count += 1
            for i in range(n - offset, start_x, -1):  # 从下至上
                nums[i][start_y] = count
                count += 1
            start_x += 1  # 更新起始点
            start_y += 1

        if n % 2 != 0:  # n为奇数时，填充中心点
            nums[mid][mid] = count
        return nums


if __name__ == "__main__":
    solution = Solution()
    matrix = solution.generateMatrix(3)
    for row in matrix:
        for item in row:
            print(item)
