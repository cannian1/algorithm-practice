# 222.完全二叉树的节点个数 https://leetcode.cn/problems/count-complete-tree-nodes/
import collections
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    """普通二叉树算节点数量(递归法)"""

    def countNodes(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        left = self.countNodes(root.left)
        right = self.countNodes(root.right)
        return left + right + 1


class Solution2:
    """迭代法"""

    def countNodes(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        queue = collections.deque([root])
        count = 0
        while queue:
            for _ in range(len(queue)):
                count += 1
                node = queue.popleft()
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
        return count


class Solution3:
    """利用完全二叉树的性质"""

    def countNodes(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        count = 0
        left = root.left
        right = root.right
        while left and right:
            count += 1
            left = left.left
            right = right.right
        if not left and not right:  # 同时到底说明这个子树是满二叉树
            return 2 ** count - 1
        return 1 + self.countNodes(root.left) + self.countNodes(root.right)
