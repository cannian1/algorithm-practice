package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 []int
		list2 []int
	}
	tests := []struct {
		name     string
		input    args
		expected []int
	}{
		{"case1", args{[]int{1, 2, 4}, []int{1, 3, 4}}, []int{1, 1, 2, 3, 4, 4}},
		{"case2", args{[]int{}, []int{0}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, ListToSlice(mergeTwoLists(SliceToList(tt.input.list1), SliceToList(tt.input.list2))),
				"mergeTwoLists(%v, %v)", tt.input.list1, tt.input.list2)
		})
	}
}
