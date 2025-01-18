package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getIntersectionNode(t *testing.T) {
	tests := []struct {
		name         string
		headA        []int
		headB        []int
		intersection []int
		expected     []int
	}{
		{
			name:         "不相交",
			headA:        []int{1, 2, 3},
			headB:        []int{4, 5, 6},
			intersection: nil,
			expected:     nil,
		},
		{
			name:         "在节点处相交",
			headA:        []int{1, 2},
			headB:        []int{3},
			intersection: []int{4, 5},
			expected:     []int{4, 5},
		},
		{
			name:         "一个链表为空",
			headA:        []int{},
			headB:        []int{1, 2, 3},
			intersection: nil,
			expected:     nil,
		},
		{
			name:         "两个链表为空",
			headA:        []int{},
			headB:        []int{},
			intersection: nil,
			expected:     nil,
		},
		{
			name:         "链表完全重叠",
			headA:        []int{1, 2},
			headB:        []int{1, 2},
			intersection: []int{1, 2},
			expected:     []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			headA := SliceToList(tt.headA)
			headB := SliceToList(tt.headB)
			intersect := SliceToList(tt.intersection)

			if intersect != nil {
				if headA != nil {
					cur := headA
					for cur.Next != nil {
						cur = cur.Next
					}
					cur.Next = intersect
				}

				if headB != nil {
					cur := headB
					for cur.Next != nil {
						cur = cur.Next
					}
					cur.Next = intersect
				}
			}

			result := getIntersectionNode(headA, headB)

			var resultSlice []int
			if result != nil {
				resultSlice = ListToSlice(result)
			}

			assert.Equal(t, tt.expected, resultSlice)
		})
	}
}
