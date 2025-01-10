// 94.二叉树的中序遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal/

package binary_tree

import (
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法中序遍历（左根右）
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)
	return slices.Concat(left, []int{root.Val}, right)
}

// 迭代法中序遍历（左根右）
func inorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	var stack []*TreeNode

	for root != nil || len(stack) > 0 {
		// 一路将当前节点的左下方向子节点全部压栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 到底了就弹出
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 采集结果，向右遍历
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
