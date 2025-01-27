// 39.组合总和
// https://leetcode.cn/problems/combination-sum/

package backtracking

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	combination := make([]int, 0)

	var dfs func(start int, target int)
	dfs = func(start int, target int) {
		if target == 0 {
			// 创建组合的副本，避免回溯过程中覆盖数据
			temp := make([]int, len(combination))
			copy(temp, combination)
			result = append(result, temp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] <= target {
				// 选择当前候选数字
				combination = append(combination, candidates[i])
				// 递归探索更新后的目标值
				dfs(i, target-candidates[i])
				// 回溯以探索其他可能性
				combination = combination[:len(combination)-1]
			}
		}
	}

	dfs(0, target)
	return result
}
