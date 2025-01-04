package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name     string
		input    args
		expected []int
	}{
		{"case1", args{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
		{"case2", args{[]int{5, 7, 7, 8, 8, 10}, 6}, []int{-1, -1}},
		{"case3", args{[]int{}, 0}, []int{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, searchRange(tt.input.nums, tt.input.target))
		})
	}
}
