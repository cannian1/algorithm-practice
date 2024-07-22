# 238. 除自身以外数组的乘积 https://leetcode.cn/problems/product-of-array-except-self/
from typing import List


class Solution:
    """左右数组 遍历三次数组，使用两个数组 3n 2n"""
    # https://leetcode.cn/problems/gou-jian-cheng-ji-shu-zu-lcof/solutions/111914/liang-tang-bian-li-by-z1m
    """
👉 正向遍历，L[i]=L[i−1]×a[i−1]
👈 反向遍历，R[j]=R[j+1]×a[j+1]
👉 正向遍历，result[i]=L[i]=L[i]×R[i]
    """

    def productExceptSelf(self, arrayA: List[int]) -> List[int]:
        if not arrayA:
            return []
        length = len(arrayA)
        L, R = [1] * length, [1] * length
        for i in range(1, length):
            L[i] = L[i - 1] * arrayA[i - 1]
        for j in range(length - 2, -1, -1):  # [倒数第二个数,0] 逆序遍历，也可以写成 for j in reversed(range(length - 1)):
            R[j] = R[j + 1] * arrayA[j + 1]
        for i in range(length):
            L[i] = L[i] * R[i]
        return L


class Solution2:
    """优化版：空间复杂度O(1)"""

    def productExceptSelf(self, arrayA: List[int]) -> List[int]:
        if not arrayA:
            return []
        length = len(arrayA)
        answer = [0] * length

        # ans[i] 表示索引 i 左侧所有元素的乘积
        # 因为索引为 '0' 的元素左侧没有元素， 所以 ans[0] = 1
        answer[0] = 1
        for i in range(1, length):
            answer[i] = arrayA[i - 1] * answer[i - 1]

        # R 为右侧所有元素的乘积
        # 刚开始右边没有元素，所以 R = 1
        R = 1
        for i in reversed(range(length)):
            # 对于索引 i，左边的乘积为 answer[i]，右边的乘积为 R
            answer[i] = answer[i] * R
            # R 需要包含右边所有的乘积，所以计算下一个结果时需要将当前值乘到 R 上
            R *= arrayA[i]

        return answer


arr = [1, 2, 3, 4, 5, 6]
s = Solution2()
res = s.productExceptSelf(arr)
print(res)
