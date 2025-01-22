// 437.路径总和 III
// https://leetcode.cn/problems/path-sum-iii/

package binary_tree

func pathSum(root *TreeNode, targetSum int) int {
	// 前缀和哈希表，用于存储到每个节点的累加和出现的次数
	prefixSum := map[int]int{0: 1} // 初始化：前缀和为 0 的路径有 1 条

	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, currSum int) int {
		if node == nil {
			return 0 // 如果节点为空，返回 0
		}

		// 更新当前路径的累加和
		currSum += node.Val
		// 计算以当前节点为结束点，满足目标和的路径数量
		count := prefixSum[currSum-targetSum]
		// 更新前缀和哈希表，记录当前路径的累加和
		prefixSum[currSum]++

		// 递归遍历左子树和右子树
		count += dfs(node.Left, currSum)
		count += dfs(node.Right, currSum)

		// 回溯：移除当前路径的累加和，防止影响其他路径
		prefixSum[currSum]--

		return count
	}

	return dfs(root, 0) // 从根节点开始递归，初始累加和为 0
}
