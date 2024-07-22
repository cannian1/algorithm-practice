# 257. 二叉树的所有路径 https://leetcode.cn/problems/binary-tree-paths/

"""前序遍历，因为要从根开始连接路径"""
from typing import Optional, List


class TreeNode(object):
    def __init__(self, val:int=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    """递归法+回溯"""

    def binaryTreePaths(self, root: Optional[TreeNode]) -> List[int]:
        """
        :type root: TreeNode
        :rtype: List[str]
        """
        result = []
        path = []
        if not root:
            return result
        self.traversal(root, path, result)
        return result

    def traversal(self, cur: Optional[TreeNode], path: List[int], result: List[int]):
        path.append(cur.val)  # 中
        if not cur.left and not cur.right:  # 到达叶子节点
            sPath = '->'.join(map(str, path))
            result.append(sPath)
            return
        if cur.left:  # 左
            self.traversal(cur.left, path, result)
            path.pop()  # 回溯
        if cur.right:  # 右
            self.traversal(cur.right, path, result)
            path.pop()  # 回溯


class Solution2:
    """递归法+隐形回溯"""

    def binaryTreePaths(self, root: TreeNode) -> List[str]:
        path = ''
        result = []
        if not root:
            return result
        self.traversal(root, path, result)
        return result

    def traversal(self, cur: TreeNode, path: str, result: List[str]) -> None:
        path += str(cur.val)
        # 若当前节点为leave，直接输出
        if not cur.left and not cur.right:
            result.append(path)

        if cur.left:
            # + '->' 是隐藏回溯
            self.traversal(cur.left, path + '->', result)

        if cur.right:
            self.traversal(cur.right, path + '->', result)


class Solution3:
    """迭代法"""

    def binaryTreePaths(self, root: TreeNode) -> List[str]:
        stack, path_st, result = [root], [str(root.val)], []
        while stack:
            cur = stack.pop()
            path = path_st.pop()
            # 如果当前节点为叶子节点，添加路径到结果中
            if not (cur.left or cur.right):
                result.append(path)
            if cur.right:
                stack.append(cur.right)
                path_st.append(path + '->' + str(cur.right.val))
            if cur.left:
                stack.append(cur.left)
                path_st.append(path + '->' + str(cur.left.val))

        return result
