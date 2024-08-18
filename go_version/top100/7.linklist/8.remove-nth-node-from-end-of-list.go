// 19. 删除链表的倒数第 N 个结点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

package linklist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := &ListNode{Next: head}
	// 必须要有 dummy 节点，否则当链表中节点数等于 n 时，将会把头节点删掉
	dummy := slow // dummy 可以维护新的头节点的关系

	for _ = range n {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
