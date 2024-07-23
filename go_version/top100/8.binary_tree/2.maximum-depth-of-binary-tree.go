package binary_tree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		result++
		for range queue {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}
