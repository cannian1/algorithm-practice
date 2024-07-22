# 33.搜索旋转排序数组
# https://leetcode.cn/problems/search-in-rotated-sorted-array/

from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        left = 0  # 二分查找左边界（左闭）
        right = len(nums) - 1  # 二分查找右边界（右闭）

        while left <= right:
            mid = left + ((right - left) >> 1)
            if nums[mid] == target:
                # 找到目标值，直接返回索引
                return mid
            if nums[left] <= nums[mid]:
                # 左半区间有序
                if nums[left] <= target < nums[mid]:
                    right = mid - 1  # 目标值落入左半区间，更新右边界
                else:
                    left = mid + 1  # 否则在右半区间查找
            else:
                # 右半区间有序
                if nums[mid] < target <= nums[right]:
                    left = mid + 1  # 目标值落入右半区间，更新左边界
                else:
                    right = mid - 1  # 否则在左半区间查找

        return -1


arr = [3, 1]
s = Solution()
t = s.search(arr, 1)
print(t)
