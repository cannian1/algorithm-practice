// 142. 环形链表 II
// https://leetcode.cn/problems/linked-list-cycle-ii/

package linklist

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			slow = head

			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}
