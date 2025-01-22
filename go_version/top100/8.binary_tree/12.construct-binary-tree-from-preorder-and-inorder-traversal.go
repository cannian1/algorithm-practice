// 105.从前序与中序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

package binary_tree

import "slices"

// 递归 O(n²), O(n²)
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	i := slices.Index(inorder, preorder[0])
	left := buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	right := buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return &TreeNode{preorder[0], left, right}
}
