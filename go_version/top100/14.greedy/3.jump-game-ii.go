// 45.跳跃游戏 II
// https://leetcode.cn/problems/jump-game-ii/

package greedy

// 给定一个长度为 n 的整数数组 nums。初始位置为 nums[0]
// 每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
// 0 <= j <= nums[i]
// i + j < n
// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

func jump(nums []int) int {
	maxFar := 0 // 目前能跳到的最远位置
	step := 0   // 跳跃次数
	end := 0    // 上次跳跃可达范围右边界（下次的最右起跳点）

	for i := range len(nums) - 1 {
		maxFar = max(maxFar, i+nums[i])
		// 到达上次跳跃能到达的右边界了
		if i == end {
			end = maxFar // 目前能跳到的最远位置变成了下次起跳位置的右边界
			step++       // 进入下一次跳跃
		}
	}
	return step
}
