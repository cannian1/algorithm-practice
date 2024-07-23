// 226.翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree/

package binary_tree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Right)
	invertTree(root.Left)
	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		for range queue {
			node := queue[0]
			queue = queue[1:]
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}
