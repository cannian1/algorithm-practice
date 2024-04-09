# 104.二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree/
import collections


class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    """后序遍历"""

    def maxdepth(self, root: TreeNode) -> int:
        if not root:
            return 0
        left = self.maxdepth(root.left)
        right = self.maxdepth(root.right)
        return max(left, right) + 1


class Solution2:
    """层序遍历"""

    def maxdepth(self, root: TreeNode) -> int:
        if not root:
            return 0
        depth = 0
        queue = collections.deque([root])
        while queue:
            level_size = len(queue)
            depth += 1
            for _ in range(level_size):
                node = queue.popleft()
                queue.append(node.left)
                queue.append(node.right)
        return depth

