// 230. 二叉搜索树中第 K 小的元素
// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/

package binary_tree

// 二叉搜索树的中序遍历是按照键增加的顺序进行的

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1] // 弹出栈顶元素
			stack = stack[:len(stack)-1]
			k--
			if k == 0 {
				return node.Val
			}
			root = node.Right
		}
	}
	return 0
}

// 递归的方式
func kthSmallest2(root *TreeNode, k int) int {
	var count int
	var result int
	found := false

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || found {
			return
		}
		inorder(node.Left)
		if found {
			return
		}
		count++
		if count == k {
			result = node.Val
			found = true
			return
		}
		inorder(node.Right)
	}

	inorder(root)
	return result
}
