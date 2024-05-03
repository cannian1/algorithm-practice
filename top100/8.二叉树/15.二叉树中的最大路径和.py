# 二叉树中的最大路径和
# https://leetcode.cn/problems/binary-tree-maximum-path-sum/
import math
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        maxSum = -math.inf

        def maxGain(node: Optional[TreeNode]) -> int:
            if not node:
                return 0

            # 递归计算左右子节点的最大贡献值
            # 只有在最大贡献值大于 0 时，才会选取对应子节点
            leftGain = max(maxGain(node.left), 0)
            rightGain = max(maxGain(node.right), 0)

            nonlocal maxSum

            # 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
            maxSum = max(maxSum, leftGain + rightGain + node.val)
            # 返回节点的最大贡献值
            return node.val + max(leftGain, rightGain)

        maxGain(root)
        return maxSum
