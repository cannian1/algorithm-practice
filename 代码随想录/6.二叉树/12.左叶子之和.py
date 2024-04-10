# 404.左叶子之和 https://leetcode.cn/problems/sum-of-left-leaves/
import collections
from typing import Optional


class TreeNode(object):
    def __init__(self, val: int = 0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    """后序遍历"""

    def sumOfLeftLeaves(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        if root.left is None and root.right is None:
            return 0

        leftValue = self.sumOfLeftLeaves(root.left)  # 左
        if root.left and not root.left.left and not root.right.right:  # 左子树是左叶子的情况
            leftValue = root.left.value

        rightValue = self.sumOfLeftLeaves(root.right)  # 右

        sum_val = leftValue + rightValue
        return sum_val


class Solution2:
    def sumOfLeftLeaves(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        sum_val = 0
        queue = collections.deque([root])
        while queue:
            node = queue.popleft()

            # 存在左子树
            if node.left:
                # 判断是否是左叶子
                if not node.left.left and not node.left.right:
                    sum_val += node.left.val
                else:
                    queue.append(node.left)
            if node.right:
                queue.append(node.right)
        return sum_val
