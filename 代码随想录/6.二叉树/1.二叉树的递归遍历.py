# 144.二叉树的前序遍历 https://leetcode.cn/problems/binary-tree-preorder-traversal
# 145.二叉树的后序遍历
# 94.二叉树的中序遍历

from typing import List


class TreeNode:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    """前序遍历（根左右）"""

    def preorderTraversal(self, root: TreeNode) -> List[int]:
        if not root:
            return []

        left = self.preorderTraversal(root.left)
        right = self.preorderTraversal(root.right)

        return [root.val] + left + right

    def inorderTraversal(self, root: TreeNode) -> List[int]:
        """中序遍历（左根右）"""
        if root is None:
            return []

        left = self.inorderTraversal(root.left)
        right = self.inorderTraversal(root.right)

        return left + [root.val] + right

    def postorderTraversal(self, root: TreeNode) -> List[int]:
        """后序遍历（左右根）"""
        if not root:
            return []

        left = self.postorderTraversal(root.left)
        right = self.postorderTraversal(root.right)

        return left + right + [root.val]


# 创建示例二叉树
root = TreeNode(1)
root.left = TreeNode(2)
root.right = TreeNode(3)
root.left.left = TreeNode(4)
root.left.right = TreeNode(5)

solution = Solution()
result = solution.preorderTraversal(root)

print(result)