// 25. K 个一组翻转链表
// https://leetcode.cn/problems/reverse-nodes-in-k-group/

package linklist

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	// 创建一个虚拟头节点，简化边界处理
	dummy := &ListNode{Next: head}
	prevGroupEnd := dummy
	cur := head

	// 计算链表的长度
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	for length >= k {
		groupStart := cur
		var prev *ListNode
		for i := 0; i < k; i++ {
			temp := cur.Next
			cur.Next = prev
			prev = cur
			cur = temp
		}

		// 将翻转后的部分连接到前后部分
		prevGroupEnd.Next = prev
		groupStart.Next = cur
		prevGroupEnd = groupStart

		// 减去已处理的 k 个节点
		length -= k
	}

	return dummy.Next
}
