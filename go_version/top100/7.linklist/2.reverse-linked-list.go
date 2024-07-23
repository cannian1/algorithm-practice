// 206. 反转链表
// https://leetcode.cn/problems/reverse-linked-list/

package linklist

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
