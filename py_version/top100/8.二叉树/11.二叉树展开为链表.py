# 114.二叉树展开为链表
# https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def flatten(self, root: Optional[TreeNode]) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        if not root:
            return
        preorder_list = []

        def preorder(root: TreeNode):
            if not root:
                return
            preorder_list.append(root)
            preorder(root.left)
            preorder(root.right)

        preorder(root)

        for i in range(1, len(preorder_list)):
            pre = preorder_list[i - 1]
            cur = preorder_list[i]
            pre.left = None
            pre.right = cur