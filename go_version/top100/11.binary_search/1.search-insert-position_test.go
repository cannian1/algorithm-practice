package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name     string
		input    args
		expected int
	}{
		{"case1", args{[]int{1, 3, 5, 6}, 5}, 2},
		{"case2", args{[]int{1, 3, 5, 6}, 2}, 1},
		{"case3", args{[]int{1, 3, 5, 6}, 7}, 4},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.expected, searchInsert(tt.input.nums, tt.input.target))
	}
}
