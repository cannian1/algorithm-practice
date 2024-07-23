package linklist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func TestGetIntersectionNode(t *testing.T) {

	commonNode := CreateListWithLoop([]int{1, 8, 4, 5})

	listA := &ListNode{Val: 4, Next: commonNode}
	listB := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: commonNode}}

	linkRes := getIntersectionNode(listA, listB)
	PrintList(linkRes)
}

func TestReverseList(t *testing.T) {
	list1 := CreateListWithLoop([]int{1, 2, 3, 4, 5})

	r1 := reverseList(list1)
	assert.Equal(t, ConvertListToSlice(r1), []int{5, 4, 3, 2, 1})
}

func TestIsPalindrome(t *testing.T) {
	list1 := CreateListWithLoop([]int{1, 2, 2, 1})
	list2 := CreateListWithLoop([]int{1, 2})

	assert.True(t, isPalindrome(list1))
	assert.False(t, isPalindrome(list2))
}
