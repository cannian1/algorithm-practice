package linklist

// ListToSlice 将链表转为切片
func ListToSlice(head *ListNode) []int {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// SliceToList 将切片转为链表
func SliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := &ListNode{}
	cur := dummy
	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return dummy.Next
}
