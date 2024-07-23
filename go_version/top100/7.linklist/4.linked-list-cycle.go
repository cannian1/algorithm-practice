// 141.环形链表
// https://leetcode.cn/problems/linked-list-cycle/

package linklist

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			return true
		}
	}
	return false
}
