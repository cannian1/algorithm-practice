// 19. 删除链表的倒数第 N 个结点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

package linklist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := &ListNode{Next: head}
	dummy := slow

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
