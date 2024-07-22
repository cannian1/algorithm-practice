# 98. 验证二叉搜索树
# https://leetcode.cn/problems/validate-binary-search-tree/
from math import inf
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    """前序"""

    def isValidBST(self, root: Optional[TreeNode], left=-inf, right=inf) -> bool:
        if not root:
            return True
        x = root.val
        return (left < x < right
                and self.isValidBST(root.left, left, x)
                and self.isValidBST(root.right, x, right))


class Solution2:
    """中序"""
    pre = -inf

    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        if not root:
            return True
        if not self.isValidBST(root.left) or root.val <= self.pre:
            return False
        self.pre = root.val
        return self.isValidBST(root.right)
