// 94.二叉树的中序遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal/

package binary_tree

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)
	return slices.Concat(left, []int{root.Val}, right)
}
