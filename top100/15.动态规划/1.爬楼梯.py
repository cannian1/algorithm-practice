# 70.爬楼梯
# https://leetcode.cn/problems/climbing-stairs/

# 一阶：1种
# 二阶：2种
# 三阶：3种 (一阶+二阶)
# 四阶：5种 （三阶+四阶）

# dp[i] 达到i阶有dp[i]种方法

# 空间复杂度为O(n)版本
class Solution:
    def climbStairs(self, n: int) -> int:
        if n <= 1:
            return n

        dp = [0] * (n + 1)
        dp[1] = 1
        dp[2] = 2

        for i in range(3, n + 1):
            dp[i] = dp[i - 1] + dp[i - 2]  # 当前阶数有几种方法依赖于前两个阶数的和

        return dp[n]


# 空间复杂度为O(3)版本
class Solution2:
    def climbStairs(self, n: int) -> int:
        if n <= 1:
            return n

        dp = [0] * 3
        dp[1] = 1
        dp[2] = 2

        for i in range(3, n + 1):
            total = dp[1] + dp[2]
            dp[1] = dp[2]
            dp[2] = total

        return dp[2]


# 空间复杂度为O(1)版本
class Solution3:
    def climbStairs(self, n: int) -> int:
        if n <= 1:
            return n

        prev1 = 1
        prev2 = 2

        for i in range(3, n + 1):
            total = prev1 + prev2
            prev1 = prev2
            prev2 = total

        return prev2
