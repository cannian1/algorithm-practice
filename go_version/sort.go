package main

import "slices"

func selectSort(nums []int) {
	n := len(nums)
	for i := range n - 1 {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}

func bubbleSort(nums []int) {
	n := len(nums)
	var flag bool
	for i := n - 1; i > 0; i-- {
		flag = false
		for j := range i {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivot := nums[0]
	var left, right []int
	for _, val := range nums {
		if val <= pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}
	return slices.Concat(quickSort(left), []int{pivot}, quickSort(right))
}

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		j := i - 1
		base := nums[i]
		for j > 0 && nums[j] > base {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = nums[j]
	}
}

// 归并排序
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) >> 1
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

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
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMid(head)
	rightSorted := sortList(mid.Next)
	mid.Next = nil
	leftSorted := sortList(head)
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
