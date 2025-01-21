//  24. 两两交换链表中的节点
// https://leetcode.cn/problems/swap-nodes-in-pairs/

package linklist

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy

	// 遍历链表
	for head != nil && head.Next != nil {
		// 定义要交换的两个节点
		first := head
		second := head.Next

		// 交换这两个节点
		pre.Next = second
		first.Next = second.Next
		second.Next = first

		// 更新 prev 和 head 指针
		pre = first
		head = first.Next
	}

	// 返回交换后的链表头节点
	return dummy.Next
}

// 递归
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}
