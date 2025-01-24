// 236. 二叉树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/

package binary_tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root // 左右都找到，当前节点是最近公共祖先
}

// 拓展
// 235. 二叉搜索树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val { // p 和 q 都在左子树
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val { // p 和 q 都在右子树
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root // 其它
}
