// 234. 回文链表
// https://leetcode.cn/problems/palindrome-linked-list/

package linklist

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 找到前半部分链表的尾节点并反转后半部分链表
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverse(firstHalfEnd.Next)

	result := true
	firstPos := head
	secondPos := secondHalfStart

	for result && secondPos != nil {
		if firstPos.Val != secondPos.Val {
			result = false
		}
		firstPos = firstPos.Next
		secondPos = secondPos.Next
	}

	firstHalfEnd.Next = reverse(secondHalfStart)
	return result
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
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
