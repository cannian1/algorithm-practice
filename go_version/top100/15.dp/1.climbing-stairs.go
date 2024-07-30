// 70.爬楼梯
// https://leetcode.cn/problems/climbing-stairs/

package dp

// 一阶：1种
// 二阶：2种
// 三阶：3种 (一阶+二阶)
// 四阶：5种 （三阶+四阶）

// dp[i] 达到i阶有dp[i]种方法

func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs2(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, 3)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		total := dp[1] + dp[2]
		dp[1] = dp[2]
		dp[2] = total
	}
	return dp[2]
}
