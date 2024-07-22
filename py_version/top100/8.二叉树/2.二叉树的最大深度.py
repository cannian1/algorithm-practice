# 104.二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree/
import collections
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    """层序遍历"""

    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0

        queue = collections.deque([root])
        max_length = 0
        while queue:
            max_length += 1
            for _ in range(len(queue)):
                cur = queue.popleft()
                if cur.left:
                    queue.append(cur.left)
                if cur.right:
                    queue.append(cur.right)
        return max_length


class Solution2:
    """后序遍历"""

    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0

        left = self.maxDepth(root.left)
        right = self.maxDepth(root.right)
        return max(left, right) + 1


class Solution3:
    """递归还可以把值传下去"""

    def maxDepth(self, root: Optional[TreeNode]) -> int:
        ans = 0

        def f(node: Optional[TreeNode], cnt: int):
            if not node:
                return
            cnt += 1
            nonlocal ans
            f(node.left, cnt)
            f(node.right, cnt)

        f(root, 0)
        return ans
