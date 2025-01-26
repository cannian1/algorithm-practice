// 46.全排列
// https://leetcode.cn/problems/permutations/

package backtracking

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案

func permute(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	backtrack(nums, used, []int{}, &res)
	return res
}

func backtrack(nums []int, used []bool, path []int, res *[][]int) {
	if len(path) == len(nums) {
		// 复制当前路径到结果中
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			// 标记当前元素为已使用
			used[i] = true
			// 递归处理下一个元素
			backtrack(nums, used, append(path, nums[i]), res)
			// 回溯，恢复当前元素为未使用
			used[i] = false
		}
	}
}
