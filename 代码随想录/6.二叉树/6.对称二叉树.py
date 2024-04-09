# 101. 对称二叉树 https://leetcode.cn/problems/symmetric-tree/
import collections


class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    """递归法，只能用后序（左右中），只有后序才可以将底部孩子的信息返回给上一层，即左孩子和右孩子的都处理完知道是对称了再处理父节点的事"""

    def isSymmetric(self, root: TreeNode) -> bool:
        if not root:
            return True
        return self.compare(root.left, root.right)

    def compare(self, left, right) -> bool:
        # 递归的终止条件
        # 首先排除空节点的情况
        if left is None and right is not None:
            return False
        elif left is not None and right is None:
            return False
        # 排除了空节点，再排除数值不相同的情况
        elif left.val != right.val:
            return False

        # 此时就是：左右节点都不为空，且数值相同的情况
        # 此时才做递归，做下一层的判断
        outside = self.compare(left.left, right.right)
        inside = self.compare(left.right, right.left)
        isSame = outside and inside
        return isSame


class Solution2:
    """迭代法（使用队列）"""

    def isSymmetric(self, root: TreeNode) -> bool:
        if not root:
            return True
        queue = collections.deque()
        queue.append(root.left)  # 将左子树头结点加入队列
        queue.append(root.right)  # 将右子树头结点加入队列

        while queue:  # 接下来就要判断这这两个树是否相互翻转
            leftNode = queue.popleft()
            rightNode = queue.popleft()
            if not leftNode and not rightNode:  # 左右结点为空，此时说明是对称的
                continue

            # 左右一个节点不为空，或者都不为空但数值不相同时，返回 false
            if not leftNode or not rightNode or leftNode.val != rightNode.val:
                return False
            queue.append(leftNode.left)  # 加入左节点左孩子
            queue.append(rightNode.right)  # 加入右节点右孩子
            queue.append(leftNode.right)  # 加入左节点右孩子
            queue.append(rightNode.left)  # 加入右节点左孩子
        return True


class Solution3:
    def isSymmetric(self, root: TreeNode) -> bool:
        if not root:
            return True
        st = []  # 这里改成了栈
        st.append(root.left)
        st.append(root.right)
        while st:
            rightNode = st.pop()
            leftNode = st.pop()
            if not leftNode and not rightNode:
                continue
            if not leftNode or not rightNode or leftNode.val != rightNode.val:
                return False
            st.append(leftNode.left)
            st.append(rightNode.right)
            st.append(leftNode.right)
            st.append(rightNode.left)
        return True


class Solution4:
    """层序遍历"""

    def isSymmetric(self, root: TreeNode) -> bool:
        if not root:
            return True

        queue = collections.deque([root.left, root.right])
        while queue:
            level_size = len(queue)

            if level_size % 2 != 0:  # 本层左右节点数量
                return False
            level_vals = []
            for i in range(level_size):
                node = queue.popleft()
                if node:
                    level_vals.append(node.val)
                    queue.append(node.left)
                    queue.append(node.right)
                else:
                    level_vals.append(None)

            if level_vals != level_vals[::-1]:
                return False
        return True
