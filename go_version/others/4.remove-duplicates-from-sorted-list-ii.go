// 82.删除排序链表中的重复元素 II
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii

package others

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		val := cur.Next.Val
		if cur.Next.Next.Val == val {
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
