# 105.从前序与中序遍历序列构造二叉树
# https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        # 递归终止条件
        if not preorder:
            return None

        # 前序遍历的第一个就是当前的中间节点
        root_val = preorder[0]
        root = TreeNode(root_val)

        # 找切割点
        idx = inorder.index(root_val)

        # 切割 inorder 数组，得到 inorder 数组的左,右半边
        inorder_left = inorder[:idx]
        inorder_right = inorder[1 + idx:]

        # 切割 preorder. 得到 preorder 数组的左,右半边
        preorder_left = preorder[1:1 + idx]
        preorder_right = preorder[1 + idx:]

        # 递归
        root.left = self.buildTree(preorder_left, inorder_left)
        root.right = self.buildTree(preorder_right, inorder_right)

        return root


class Solution2:
    """用哈希表优化时间复杂度"""

    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        # 建立哈希表，存储中序遍历的值到索引的映射
        inorder_index_map = {value: index for index, value in enumerate(inorder)}

        def array_to_tree(left, right):
            if left > right:
                return None

            # 前序遍历的第一个元素是根节点(每轮递归都会创建一个新的 preorder，元素为上一轮移除头节点的 preorder)
            root_val = preorder.pop(0)
            root = TreeNode(root_val)

            # 划分左右子树
            index = inorder_index_map[root_val]

            # 递归构建左右子树
            root.left = array_to_tree(left, index - 1)
            root.right = array_to_tree(index + 1, right)

            return root

        return array_to_tree(0, len(inorder) - 1)


preorder = [3, 9, 20, 15, 7]
inorder = [9, 3, 15, 20, 7]

s = Solution2()
s.buildTree(preorder, inorder)
