// 114.二叉树展开为链表
// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list

package binary_tree

// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。

func flatten(root *TreeNode) {
	var head *TreeNode
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		dfs(node.Left)
		node.Left = nil
		node.Right = head // 头插法，相当于链表的 node.Next = head
		head = node       // 现在链表头节点是 node
	}
	dfs(root)
}
