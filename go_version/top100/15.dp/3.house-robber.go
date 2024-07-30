// 198.打家劫舍
// https://leetcode.cn/problems/house-robber/

package dp

// 最多可以偷窃的金额为 dp[i]
// 不打劫当前房子，金额为 dp[i-1]; 打劫当前房子，金额为 dp[i-2] + nums[i-1]

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)

	dp[1] = nums[0] // 第一个房子的最大打劫金额

	for i := 2; i <= n; i++ {
		// 其中 dp[i-1] 表示不打劫当前房子，dp[i-2] + nums[i-1] 表示打劫当前房子
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[n] // 打劫到最后一个房子的最大金额
}

/*
假设输入为 nums := []int{2, 7, 9, 3, 1}：

初始化 dp 数组：dp := [0, 2, 0, 0, 0, 0]。
计算 dp[2]：dp[2] = max(dp[1], dp[0] + nums[1]) = max(2, 0 + 7) = 7。
计算 dp[3]：dp[3] = max(dp[2], dp[1] + nums[2]) = max(7, 2 + 9) = 11。
计算 dp[4]：dp[4] = max(dp[3], dp[2] + nums[3]) = max(11, 7 + 3) = 11。
计算 dp[5]：dp[5] = max(dp[4], dp[3] + nums[4]) = max(11, 11 + 1) = 12。
最终返回 dp[5] = 12，即打劫到最后一个房子的最大金额是 12。
*/
