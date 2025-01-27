// 46.全排列
// https://leetcode.cn/problems/permutations/

package backtracking

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案

func permute(nums []int) [][]int {
	result := [][]int{}

	// swap 交换两个元素的位置
	swap := func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	}

	// backtrack 实现全排列的回溯逻辑
	var backtrack func(int)
	backtrack = func(start int) {
		if start == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			result = append(result, temp)
			return
		}

		for i := start; i < len(nums); i++ {
			swap(start, i)       // 选择当前元素
			backtrack(start + 1) // 递归处理剩余元素
			swap(start, i)       // 回溯，撤销选择
		}
	}

	backtrack(0)
	return result
}

func permute2(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	backtrack(nums, used, []int{}, &res)
	return res
}

/*
开始: path = [], used = [F,F,F]
├─ 选1: path = [1], used = [T,F,F]
│   ├─ 选2: path = [1,2], used = [T,T,F]
│   │   └─ 选3: path = [1,2,3] → 存入结果 ✅
│   └─ 选3: path = [1,3], used = [T,F,T]
│       └─ 选2: path = [1,3,2] → 存入结果 ✅
├─ 选2: path = [2], used = [F,T,F]
│   ├─ 选1: path = [2,1], used = [T,T,F]
│   │   └─ 选3: path = [2,1,3] → 存入结果 ✅
│   └─ 选3: path = [2,3], used = [F,T,T]
│       └─ 选1: path = [2,3,1] → 存入结果 ✅
└─ 选3: path = [3], used = [F,F,T]
    ├─ 选1: path = [3,1], used = [T,F,T]
    │   └─ 选2: path = [3,1,2] → 存入结果 ✅
    └─ 选2: path = [3,2], used = [F,T,T]
        └─ 选1: path = [3,2,1] → 存入结果 ✅
*/

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
			// append 会生成一个新切片，原 path 保持不变，就不需要 pop 了
			backtrack(nums, used, append(path, nums[i]), res)
			// 回溯，恢复当前元素为未使用
			used[i] = false
		}
	}
}
