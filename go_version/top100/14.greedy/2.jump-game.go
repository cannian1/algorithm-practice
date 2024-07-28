// 55.跳跃游戏
// https://leetcode.cn/problems/jump-game/

package greedy

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
