package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name     string
		input    args
		expected []int
	}{
		{"case1", args{[]int{2, 4, 3}, []int{5, 6, 4}}, []int{7, 0, 8}},
		{"case2", args{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
		{"case3", args{[]int{0}, []int{0}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, ListToSlice(addTwoNumbers(SliceToList(tt.input.l1), SliceToList(tt.input.l2))),
				"addTwoNumbers(%v, %v)", tt.input.l1, tt.input.l2)
		})
	}
}
