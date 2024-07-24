// 199.二叉树的右视图
// https://leetcode.cn/problems/binary-tree-right-side-view/

package binary_tree

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// 在开始遍历层之前，先确定这一层的节点数量
		levelSize := len(queue)
		result = append(result, queue[levelSize-1].Val)

		for _ = range levelSize {
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
