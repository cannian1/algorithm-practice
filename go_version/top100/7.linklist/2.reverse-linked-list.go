// 206. 反转链表
// https://leetcode.cn/problems/reverse-linked-list/

package linklist

// 迭代法
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head

	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

// 递归法
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
