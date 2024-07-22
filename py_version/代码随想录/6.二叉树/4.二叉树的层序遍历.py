# 102.二叉树的层序遍历
# https://leetcode.cn/problems/binary-tree-level-order-traversal/
# 107.二叉树的层次遍历 II https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
# 199.二叉树的右视图 https://leetcode.cn/problems/binary-tree-right-side-view/
# 637.二叉树的层平均值 https://leetcode.cn/problems/average-of-levels-in-binary-tree/
# 429.N叉树的层序遍历 https://leetcode.cn/problems/n-ary-tree-level-order-traversal/
# 515.在每个树行中找最大值 https://leetcode.cn/problems/find-largest-value-in-each-tree-row/
# 116.填充每个节点的下一个右侧节点指针 https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/
# 104.二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree/
# 111.二叉树的最小深度 https://leetcode.cn/problems/minimum-depth-of-binary-tree/
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
            level_size = len(queue)
            for _ in range(level_size):  # 遍历每层的节点
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
            level_size = len(queue)
            for _ in range(level_size):
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


class Solution5:
    class Node:
        def __init__(self, val=None, children=None):
            self.val = val
            self.children = children

    def levelOrder(self, root: Node) -> List[List[int]]:
        """429.N叉树的层序遍历"""
        if not root:
            return []

        result = []
        queue = collections.deque([root])

        while queue:
            level_size = len(queue)
            level = []

            for _ in range(level_size):
                node = queue.popleft()
                level.append(node.val)

                for child in node.children:
                    queue.append(child)
                result.append(level)

        return result


class Solution6:
    def largestValues(self, root: TreeNode) -> List[int]:
        # 515.在每个树行中找最大值 https://leetcode.cn/problems/find-largest-value-in-each-tree-row/
        if not root:
            return []

        result = []
        queue = collections.deque([root])
        while queue:
            max_num = float('-inf')
            level_size = len(queue)
            for _ in range(level_size):
                node = queue.popleft()
                max_num = max(max_num, node.val)

                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            result.append(max_num)
        return result


class Solution7:
    class Node:
        def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
            self.val = val
            self.left = left
            self.right = right
            self.next = next

    def connect(self, root: 'Node') -> 'Node':
        # 116.填充每个节点的下一个右侧节点指针 https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/
        if not root:
            return root

        queue = collections.deque([root])
        # 单层遍历的时候记录一下本层的头部节点，然后在遍历的时候让前一个节点指向本节点就可以了
        while queue:
            level_size = len(queue)
            prev = None

            for _ in range(level_size):
                node = queue.popleft()

                if prev:
                    prev.next = node
                prev = node

                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)

        return root


class Solution8:
    def maxDepth(self, root: TreeNode) -> int:
        # 104.二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree/
        depth = 0
        if not root:
            return depth

        queue = collections.deque([root])
        while queue:
            depth += 1  # 这里是每层，下面的 for 循环是每层的每个节点
            level_len = len(queue)
            for _ in range(level_len):
                node = queue.popleft()
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)

        return depth


class Solution9:
    """只有当左右孩子都为空的时候，才说明遍历的最低点了。如果其中一个孩子为空则不是最低点"""

    def maxDepth(self, root: TreeNode) -> int:
        # 111.二叉树的最小深度 https://leetcode.cn/problems/minimum-depth-of-binary-tree/
        depth = 0
        if not root:
            return depth

        queue = collections.deque([root])
        while queue:
            depth += 1
            level_size = len(queue)
            for _ in range(level_size):
                node = queue.popleft()

                if not node.left and not node.right:
                    return depth

                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
        return depth
