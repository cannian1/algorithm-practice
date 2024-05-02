# 543.二叉树的直径
# https://leetcode.cn/problems/diameter-of-binary-tree/
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


"""两个叶子节点之间路径 = 根节点左右子节点深度之和"""


class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        ans = 0

        def dfs(node: Optional[TreeNode]):
            if not node:
                return 0
            # 左右子树“往下”的最大路径长度
            left = dfs(node.left)
            right = dfs(node.right)
            nonlocal ans
            # 用左右子树为根往下最大路径和的最大值更新最大直径
            ans = max(ans, left + right)
            return max(left, right) + 1  # 在子路径基础上增加当前节点

        dfs(root)
        return ans
