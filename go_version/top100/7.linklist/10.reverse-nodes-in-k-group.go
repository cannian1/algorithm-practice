// 25. K 个一组翻转链表
// https://leetcode.cn/problems/reverse-nodes-in-k-group/

package linklist

func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	cur := head

	for cur != nil {
		n += 1
		cur = cur.Next
	}

	dummy := &ListNode{Next: head}
	p0 := dummy
	var pre *ListNode = nil
	cur = head

	for n >= k {
		n -= k
		for _ = range k {
			temp := cur.Next
			cur.Next = pre
			pre = cur
			cur = temp
		}

		nxt := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = nxt
	}
	return dummy.Next
}
