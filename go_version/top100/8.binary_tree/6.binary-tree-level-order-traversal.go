// 102.二叉树的层序遍历
// https://leetcode.cn/problems/binary-tree-level-order-traversal/

package binary_tree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := make([]int, 0)
		for range queue {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
