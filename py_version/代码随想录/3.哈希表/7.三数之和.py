# 15. 三数之和
# https://leetcode.cn/problems/3sum/


from typing import List

"""
不可以包含重复的三元组!!!
涉及去重考虑的条件太多，不适合使用哈希表解决

好的解法是使用双指针
nums 口 口 口 口 口 口 口
     i left         right
 
如果 nums[i] + nums[left] + nums[right] > 0 
    right 左移
如果 nums[i] + nums[left] + nums[right] < 0 
    left 右移
将结果集存放到一个二维数组里
关键：去重
"""


class Solution:
    """双指针"""

    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums = sorted(nums)
        result = []

        for i in range(len(nums)):
            # 如果第一个元素已经大于 0 ，不需要继续检查
            if nums[i] > 0:
                return result

            # 跳过相同的元素避免重复
            # nums[i] 要与 nums[i-1]比较而不能与 nums[i+1]比较
            #
            if i > 0 and nums[i] == nums[i - 1]:
                continue

            left = i + 1
            right = len(nums) - 1

            while left < right:
                sum_ = nums[i] + nums[left] + nums[right]
                if sum_ < 0:
                    left += 1
                elif sum_ > 0:
                    right -= 1
                else:
                    result.append([nums[i], nums[left], nums[right]])

                    # 跳过相同的元素以避免重复
                    while right > left and nums[right] == nums[right - 1]:
                        right -= 1
                    while right > left and nums[left] == nums[left + 1]:
                        left += 1
                    right -= 1
                    left += 1
        return result
