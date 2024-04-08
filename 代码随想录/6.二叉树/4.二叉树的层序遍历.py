# 102.二叉树的层序遍历
# https://leetcode.cn/problems/binary-tree-level-order-traversal/
# 107.二叉树的层次遍历 II https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
# 199.二叉树的右视图 https://leetcode.cn/problems/binary-tree-right-side-view/
# 637.二叉树的层平均值 https://leetcode.cn/problems/average-of-levels-in-binary-tree/
import collections
from typing import Optional, List


class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution(object):
    """102.二叉树层序遍历"""

    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        if not root:
            return []
        queue = collections.deque([root])  # 创建一个队列用于保存遍历的元素
        result = []
        while queue:
            level = []  # 用于保存每层的队列元素的值
            for _ in range(len(queue)):  # 遍历每层的节点
                cur = queue.popleft()  # 队尾元素出队
                level.append(cur.val)
                if cur.left:
                    queue.append(cur.left)  # 将子节点入队
                if cur.right:
                    queue.append(cur.right)
            result.append(level)  # 将当前层的元素值收集到返回结果中
        return result


class Solution2:
    # 107.二叉树的层次遍历 II (最后一步将数组逆序就是自底向上层次遍历了)
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if not root:
            return []

        queue = collections.deque([root])
        result = []
        while queue:
            level = []
            for _ in range(len(queue)):
                cur = queue.popleft()
                level.append(cur.val)
                if cur.left:
                    queue.append(cur.left)
                if cur.right:
                    queue.append(cur.right)
            result.append(level)
        return result[::-1]


class Solution3:
    """199.二叉树的右视图"""

    def rightSideView(self, root: TreeNode) -> List[int]:
        if not root:
            return []

        queue = collections.deque([root])
        right_view = []

        while queue:
            level_size = len(queue)
            for i in range(level_size):
                node = queue.popleft()

                # 遍历到这层的最后一个节点的时候加入队列
                if i == level_size - 1:
                    right_view.append(node.val)

                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
        return right_view


class Solution4:
    """637.二叉树的层平均值"""

    def averageOfLevels(self, root: TreeNode) -> List[float]:
        if not root:
            return []

        queue = collections.deque([root])
        averages = []

        while queue:
            size = len(queue)
            level_sum = 0

            for i in range(size):
                node = queue.popleft()

                level_sum += node.val

                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)

            averages.append(level_sum / size)

        return averages
