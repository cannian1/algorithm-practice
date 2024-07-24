// 543.二叉树的直径
// https://leetcode.cn/problems/diameter-of-binary-tree/

package binary_tree

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1 // 下面 +1 后，对于叶子结点就刚好是 0
		}
		lLen := dfs(node.Left) + 1      // 左子树最大链长+1
		rLen := dfs(node.Right) + 1     // 右子树最大链长+1
		result = max(result, lLen+rLen) // 两条链拼成路径
		return max(lLen, rLen)          // 当前子树最大链长
	}
	dfs(root)
	return result
}
