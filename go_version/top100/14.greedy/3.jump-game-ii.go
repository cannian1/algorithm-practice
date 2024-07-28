// 45.跳跃游戏 II
// https://leetcode.cn/problems/jump-game-ii/

package greedy

func jump(nums []int) int {
	maxFar := 0 // 目前能跳到的最远位置
	step := 0   // 跳跃次数
	end := 0    // 上次跳跃可达范围右边界（下次的最右起跳点）

	for i := range len(nums) - 1 {
		maxFar = max(maxFar, i+nums[i])
		// 到达上次跳跃能到达的右边界了
		if i == end {
			end = maxFar // 目前能跳到的最远位置变成了下次起跳位置的有边界
			step++       // 进入下一次跳跃
		}
	}
	return step
}
