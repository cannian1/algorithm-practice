# 230. 二叉搜索树中第 K 小的元素
# https://leetcode.cn/problems/kth-smallest-element-in-a-bst/
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    """中序遍历"""

    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        ans = 0

        def bfs(root):
            nonlocal k, ans
            if not root:
                return
            bfs(root.left)
            k -= 1
            if k == 0:
                ans = root.val
                return
            bfs(root.right)

        bfs(root)
        return ans


class Solution2:
    """中序遍历(另一种写法)"""

    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        stack = []
        while root or stack:
            while root:
                stack.append(root)
                root = root.left
            root = stack.pop()
            k -= 1
            if k == 0:
                return root.val
            root = root.right
