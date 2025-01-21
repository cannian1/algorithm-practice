package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeKLists(t *testing.T) {

	tests := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{"case1", [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{"case2", [][]int{{-10, -5, 0}, {1, 2}, {0, 4, 5}}, []int{-10, -5, 0, 0, 1, 2, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lists []*ListNode
			for _, l := range tt.input {
				lists = append(lists, SliceToList(l))
			}

			result := mergeKLists(lists)
			resultSlice := ListToSlice(result)
			assert.Equalf(t, tt.expected, resultSlice, "mergeKLists(%v)", tt.input)
		})
	}
}
