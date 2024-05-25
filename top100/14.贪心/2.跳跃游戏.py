# 55.跳跃游戏
# https://leetcode.cn/problems/jump-game/
from typing import List


class Solution:
    def canJump(self, nums: List[int]) -> bool:
        cover = 0  # 覆盖范围
        if len(nums) == 1:  # 只有一个元素，就是能到达
            return True
        for i in range(len(nums)):
            if i <= cover:
                cover = max(i + nums[i], cover)
                if cover >= len(nums) - 1:
                    return True
        return False
