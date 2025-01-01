package sort

// 归并排序
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) >> 1
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

// 合并两个有序数组
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 将剩余的元素追加到结果数组中
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// 归并排序链表版
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMid(head)
	rightSorted := SortList(mid.Next)
	mid.Next = nil
	leftSorted := SortList(head)
	return mergeTwoList(leftSorted, rightSorted)
}

func getMid(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func mergeTwoList(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
