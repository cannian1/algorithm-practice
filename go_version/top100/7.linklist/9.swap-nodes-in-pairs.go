//  24. 两两交换链表中的节点
// https://leetcode.cn/problems/swap-nodes-in-pairs/

package linklist

func swapPairs(head *ListNode) *ListNode {
	cur := &ListNode{Next: head}
	dummy := cur

	for cur.Next != nil && cur.Next.Next != nil {
		temp := cur.Next
		temp1 := cur.Next.Next.Next

		cur.Next = cur.Next.Next
		cur.Next.Next = temp
		temp.Next = temp1
		cur = cur.Next.Next
	}
	return dummy.Next
}
