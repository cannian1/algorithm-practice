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

	// https://leetcode.cn/problems/reverse-linked-list/solutions/2361282/206-fan-zhuan-lian-biao-shuang-zhi-zhen-r1jel
	// 迭代法局部反转后，pre 指向本组反转后的头节点，
	// cur 指向反转前的本组最后一个节点的下一个节点，即下一组的头节点
	for length >= k {
		groupStart := cur // 保存当前组的头节点
		var prev *ListNode
		for i := 0; i < k; i++ {
			temp := cur.Next
			cur.Next = prev
			prev = cur
			cur = temp
		}

		// 将翻转后的部分连接到前后部分
		prevGroupEnd.Next = prev // prevGroupEnd 没有参与反转，仍然是上一组的尾结点，与参与反转的组连接
		groupStart.Next = cur    // 组内反转后，groupStart 是组内的尾节点
		prevGroupEnd = groupStart

		// 减去已处理的 k 个节点
		length -= k
	}

	return dummy.Next
}

// 递归
func reverseKGroup1(head *ListNode, k int) *ListNode {
	cur := head
	for range k {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	newHead := reverseFromStartToEnd(head, cur)
	head.Next = reverseKGroup1(cur, k)
	return newHead
}

func reverseFromStartToEnd(start, end *ListNode) *ListNode {
	var pre *ListNode
	cur := start
	for cur != end {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
