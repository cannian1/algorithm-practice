package ds

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateListWithLoop 使用for循环创建单链表
func CreateListWithLoop(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	cur := head

	for i := 1; i < len(values); i++ {
		cur.Next = &ListNode{Val: values[i]}
		cur = cur.Next
	}

	return head
}

func ConvertListToSlice(head *ListNode) []int {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func PrintList(head *ListNode) {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	fmt.Println(result)
}
