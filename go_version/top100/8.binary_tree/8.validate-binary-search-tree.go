// 98. 验证二叉搜索树
// https://leetcode.cn/problems/validate-binary-search-tree/

package binary_tree

import "math"

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, left, right int) bool
	dfs = func(node *TreeNode, left, right int) bool {
		if node == nil {
			return true
		}

		x := node.Val
		return left < x && x < right && dfs(node.Left, left, x) && dfs(node.Right, x, right)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}
