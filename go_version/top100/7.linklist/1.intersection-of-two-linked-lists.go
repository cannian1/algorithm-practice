// 160. 相交链表
// https://leetcode.cn/problems/intersection-of-two-linked-lists/

package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	dis := getLength(headA) - getLength(headB)
	if dis > 0 {
		headA = moveForward(headA, dis)
	} else {
		headB = moveForward(headB, -dis)
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		head = head.Next
		length++
	}
	return length
}

func moveForward(head *ListNode, steps int) *ListNode {
	for steps > 0 {
		head = head.Next
		steps--
	}
	return head
}
