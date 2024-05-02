# 108. 将有序数组转换为二叉搜索树
# https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/
from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        return self.build(nums, 0, len(nums) - 1)

    def build(self, nums, left, right):
        if left > right:
            return None
        # 以数组中间位置的元素作为根
        mid = left + ((right - left) >> 1)
        ans = TreeNode(nums[mid])
        ans.left = self.build(nums, left, mid - 1)
        ans.right = self.build(nums, mid + 1, right)
        return ans


arr = [-10,-3,0,5,9]
s = Solution()
tr = s.sortedArrayToBST(arr)