# 15. 三数之和
# https://leetcode.cn/problems/3sum/
from typing import List


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()

        result = []

        for i in range(len(nums)):
            # 如果第一个元素已经大于 0 ，不需要继续检查
            if nums[i] > 0:
                return result

            # 跳过重复元素
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

                    # 跳过相同元素以避免重复
                    while right > left and nums[right] == nums[right - 1]:
                        right -= 1
                    while right > left and nums[left] == nums[left + 1]:
                        left += 1
                    right -= 1
                    left += 1

        return result
