// 55.跳跃游戏
// https://leetcode.cn/problems/jump-game/

package greedy

/*
给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
*/

func canJump(nums []int) bool {
	mx := 0
	for i, jmp := range nums {
		if i > mx {
			return false // 无法到达 i
		}
		mx = max(mx, i+jmp) // 从 i 最右可以跳到 i + jmp
	}
	return true
}
