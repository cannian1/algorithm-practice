from typing import List


class TreeNode:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


"""
处理：将元素放进result数组中
访问：遍历节点
"""

"""
前序遍历是中左右，每次先处理的是中间节点，那么先将根节点放入栈中，然后将右孩子加入栈，再加入左孩子
按 根右左 的顺序入栈，才会按 根左右 的顺序出栈
"""


# 前序遍历-迭代-LC144_二叉树的前序遍历
class Solution:
    def preorderTraversal(self, root: TreeNode) -> List[int]:
        """前序遍历（左中右）"""
        # 根结点为空则返回空列表
        if not root:
            return []
        stack = [root]
        result = []
        while stack:
            node = stack.pop()
            # 中结点先处理
            result.append(node.val)
            # 右孩子先入栈
            if node.right:
                stack.append(node.right)
            # 左孩子后入栈
            if node.left:
                stack.append(node.left)
        return result

    def inorderTraversal(self, root: TreeNode) -> List[int]:
        """中序遍历(中左右)"""
        if not root:
            return []
        stack = []  # 不能提前将root结点加入stack中
        result = []
        cur = root  # 指针来遍历元素，栈来保存元素
        while cur or stack:
            # 先迭代访问最底层的左子树结点
            if cur:  # 一路保存所有的左孩子，加入待处理栈中
                stack.append(cur)
                cur = cur.left
            # 到达最左结点后处理栈顶结点
            else:
                cur = stack.pop()  # 把待处理栈中的元素弹出后，可以加入返回结果序列中
                result.append(cur.val)
                # 取栈顶元素右结点
                cur = cur.right
        return result

    def postorderTraversal(self, root: TreeNode) -> List[int]:
        """后序遍历（左右中） 前序改一改顺序，最后翻转一下就可以了"""
        if not root:
            return []
        stack = [root]
        result = []
        while stack:
            node = stack.pop()
            # 中结点先处理
            result.append(node.val)
            # 左孩子先入栈
            if node.left:
                stack.append(node.left)
            # 右孩子后入栈
            if node.right:
                stack.append(node.right)
        # 将最终的数组翻转
        return result[::-1]
