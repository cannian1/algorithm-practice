# 226.翻转二叉树 https://leetcode.cn/problems/invert-binary-tree/
import collections


class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


"""可以前序和后序，中序的参数两个都要改成 root.left"""


class Solution(object):
    """递归法，前序遍历"""

    def invertTree(self, root: TreeNode) -> TreeNode:
        if not root:
            return None
        root.left, root.right = root.right, root.left
        self.invertTree(root.left)
        self.invertTree(root.right)
        return root


class Solution2:
    """层序遍历"""

    def invertTree(self, root: TreeNode) -> TreeNode:
        if not root:
            return None
        queue = collections.deque([root])

        while queue:
            for _ in range(len(queue)):
                node = queue.popleft()
                node.left, node.right = node.right, node.left
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
        return root
