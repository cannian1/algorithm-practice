# LeetCode 209. 长度最小的子数组
# https://leetcode.cn/problems/minimum-size-subarray-sum/
# target 是正整数

from typing import List


# 滑动窗口
class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:

        left = 0
        min_len = float('inf')
        cur_sum = 0  # 当前的累加值

        # 右指针先走
        for right, val in enumerate(nums):
            # 计算当前累加和
            cur_sum += val

            # 这个循环执行结束后可以得到当前右指针往左的长度最小的子数组长度
            while cur_sum >= target:  # 当前累加值大于目标值
                # 更新最短子数组长度
                min_len = min(min_len, right - left + 1)
                cur_sum -= nums[left]  # 减去最左指针当前指向元素
                left += 1  # 左指针向右移动

            right += 1

        return min_len if min_len != float('inf') else 0


# 暴力法
class SlowSolution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        l = len(nums)
        min_len = float('inf')

        for i in range(l):
            cur_sum = 0
            for j in range(i, l):
                cur_sum += nums[j]
                if cur_sum >= target:
                    min_len = min(min_len, j - i + 1)
                    break

        return min_len if min_len != float('inf') else 0


if __name__ == "__main__":
    solution = Solution()

    # 测试用例1：正常情况
    print(solution.minSubArrayLen(7, [2, 3, 1, 2, 4, 3]))  # 预期输出: 2 (因为子数组[4, 3]的和为7，是最短的满足条件的子数组)

    # 测试用例2：边界情况，最短子数组为整个数组
    print(solution.minSubArrayLen(11, [1, 2, 3, 4, 5]))  # 预期输出: 5

    # 测试用例3：异常情况，没有满足条件的子数组
    print(solution.minSubArrayLen(100, [1, 2, 3, 4, 5]))  # 预期输出: 0

    # 测试用例5：目标和为0
    print(solution.minSubArrayLen(0, [1, 2, 3, 4, 5]))  # 预期输出: 0
