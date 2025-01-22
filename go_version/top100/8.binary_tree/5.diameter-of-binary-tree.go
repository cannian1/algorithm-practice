// 543.二叉树的直径
// https://leetcode.cn/problems/diameter-of-binary-tree/

package binary_tree

// 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0

	// 枚举每个 node，假设直径在这里「拐弯」，也就是计算由左右两条从下面的叶子节点到 node 的链的节点值之和，去更新答案的最大值
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
