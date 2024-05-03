# 二叉树的最近公共祖先
# https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

"""
求最小公共祖先，需要从底向上遍历，那么二叉树，只能通过后序遍历（即：回溯）实现从底向上的遍历方式。
在回溯的过程中，必然要遍历整棵二叉树，即使已经找到结果了，依然要把其他节点遍历完，因为要使用递归函数的返回值（也就是代码中的left和right）做逻辑判断
要理解如果返回值left为空，right不为空为什么要返回right，为什么可以用返回right传给上一层结果
"""

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if root == q or root == p or not root:
            return root

        left = self.lowestCommonAncestor(root.left,p,q)
        right = self.lowestCommonAncestor(root.right,p,q)

        if left is not None and right is not None:
            return root

        if left is None and right is not None:
            return right
        elif left is not None and right is None:
            return left
        else:
            return None
