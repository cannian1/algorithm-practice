# 437.路径总和 III
# https://leetcode.cn/problems/path-sum-iii/
import collections
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    """dfs O(n²) O(n)"""

    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> int:
        def rootSum(root: Optional[TreeNode], targetSum: int) -> int:
            if not root:
                return 0

            res = 0
            # 判断当前节点的值是否等于targetSum。
            # 如果相等，说明从根节点到当前节点的路径和为targetSum，因此结果加1
            if root.val == targetSum:
                res += 1

            # 统计左右子树中符合路径和为目标值的路径数
            res += rootSum(root.left, targetSum - root.val)
            res += rootSum(root.right, targetSum - root.val)
            return res

        if not root:
            return 0

        ret = rootSum(root, targetSum)
        ret += self.pathSum(root.left, targetSum)
        ret += self.pathSum(root.right, targetSum)
        return ret


class Solution2:
    """前缀和 O(n) O(n)"""
    """两节点之间路径和即两节点的前缀和之差"""

    def pathSum(self, root: TreeNode, targetSum: int) -> int:
        prefix = collections.defaultdict(int)
        prefix[0] = 1 # 根节点到根节点的路径和为0

        def dfs(root, cur):
            if not root:
                return 0

            ret = 0
            cur += root.val
            # 检查cur - targetSum的前缀和是否在prefix中，
            # 如果在，说明存在一条从某个节点到当前节点的路径和为targetSum，将这个次数加到返回值ret上
            ret += prefix[cur - targetSum]
            prefix[cur] += 1
            ret += dfs(root.left, cur)
            ret += dfs(root.right, cur)
            # 在返回之前，将当前的前缀和cur的出现次数减1，因为在回溯过程中，当前节点不再是路径的一部分
            prefix[cur] -= 1

            return ret

        return dfs(root, 0)
