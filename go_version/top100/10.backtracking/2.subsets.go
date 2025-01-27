// 78.子集
// https://leetcode.cn/problems/subsets/

package backtracking

func subsets(nums []int) [][]int {
	result := [][]int{}
	subset := []int{}

	var dfs func(int)
	// 深度优先搜索 (DFS)
	dfs = func(index int) {
		// 将当前子集加入结果集
		temp := make([]int, len(subset))
		copy(temp, subset)
		result = append(result, temp)

		// 遍历数组中剩下的元素
		for i := index; i < len(nums); i++ {
			// 选择当前元素
			subset = append(subset, nums[i])
			// 递归调用，处理下一个元素
			dfs(i + 1)
			// 回溯，移除当前元素
			subset = subset[:len(subset)-1]
		}
	}

	dfs(0)
	return result
}
