package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binarySearchFirst(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name     string
		input    args
		expected int
	}{
		{"case1", args{[]int{1, 3, 4, 5, 6, 6, 6, 8, 8, 8, 9}, 6}, 4},
		{"case2", args{[]int{1, 3, 4, 5, 6, 6, 6, 8, 8, 8, 9}, 7}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, binarySearchFirst(tt.input.nums, tt.input.target), "binarySearchFirst(%v, %v)", tt.input.nums, tt.input.target)
		})
	}
}
