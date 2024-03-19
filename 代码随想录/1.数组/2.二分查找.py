# 704. 二分查找
# https://leetcode.cn/problems/binary-search/

class Solution:
    def search1(self, nums: list[int], target: int) -> int:
        """左闭右闭"""
        left, right = 0, len(nums) - 1  # 定义target在左闭右闭的区间里，[left, right]
        # 在这个区间中小于等于是合法的
        while left <= right:
            mid = left + (right - left) // 2  # 向下取整
            # 已经判断过 left <= right ，因此 nums[mid] 不是要搜索的值
            if nums[mid] > target:
                right = mid - 1  # target在左区间，所以 左区间的边界更新为 [left, middle - 1]
            elif nums[mid] < target:
                left = mid + 1  # target在右区间，所以 右区间的边界更新为 [middle + 1, right]
            else:
                return mid
        return -1  # 找不到目标值

    def search2(self, nums: list[int], target: int) -> int:
        """左闭右开"""
        left, right = 0, len(nums)  # 定义 target 在左闭右开区间,即：[left, right)

        while left < right:  # left == right的时候，在[left, right)是无效的空间，所以使用 <
            mid = left + (right - left) // 2

            if nums[mid] > target:
                # 下一个搜索的左区间不包含 nums[mid]，所以右边界变为mid
                right = mid  # target 在左区间，在[left, middle)中
            elif nums[mid] < target:  # 左闭，
                left = mid + 1  # target 在右区间，在[middle + 1, right)中
            else:
                return mid
        return -1
