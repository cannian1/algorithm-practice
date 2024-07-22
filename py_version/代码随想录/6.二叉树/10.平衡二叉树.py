# 110.平衡二叉树
# https://leetcode.cn/problems/balanced-binary-tree/

"""平衡二叉树：二叉树的任何一个节点左右子树高度差小于等于1"""
"""求深度可以用层序，求高度不可以"""


class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution(object):
    def isBalanced(self, root: TreeNode) -> bool:
        return self.get_height(root) != -1

    def get_height(self, root: TreeNode) -> int:
        if not root:
            return 0
        left = self.get_height(root.left)
        right = self.get_height(root.right)
        if left == -1 or right == -1 or abs(left - right) > 1:
            return -1
        return max(left, right) + 1


class Solution2(object):
    def isBalanced(self, root: TreeNode) -> bool:
        return self.get_height(root) != -1

    def get_height(self, root: TreeNode) -> int:
        if not root:
            return 0
        # 左
        if (left_height := self.get_height(root.left)) == -1:
            return -1
        # 右
        if (right_height := self.get_height(root.right)) == -1:
            return -1
        # 中
        if abs(left_height - right_height) > 1:
            return -1
        else:
            return 1 + max(left_height, right_height)



