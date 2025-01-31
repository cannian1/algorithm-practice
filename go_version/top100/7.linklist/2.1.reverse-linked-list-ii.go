// 92.反转链表 II
// https://leetcode.cn/problems/reverse-linked-list-ii/

package linklist

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prevEnd := dummy

	for range left - 1 {
		prevEnd = prevEnd.Next
	}

	var pre *ListNode
	cur := prevEnd.Next

	for range right - left + 1 {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	// 分割点没切断，还能索引到局部反转前分割点的下一个节点
	// 在反转后它已经到了末尾，迭代反转结束后的 cur 位置在下一段的起始，pre 在目标段的起始
	prevEnd.Next.Next = cur
	prevEnd.Next = pre
	return dummy.Next
}
