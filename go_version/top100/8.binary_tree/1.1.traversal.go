package binary_tree

import "slices"

// 先序遍历（根左右）
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)
	return slices.Concat([]int{root.Val}, left, right)
}

func preorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}

// 后序遍历是压栈的顺序改一下，最后一步反转切片再返回
// 即将（根右左）-> （左右根）
func postorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, node.Val)

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	reverse(res)
	return res
}

func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left <= right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
