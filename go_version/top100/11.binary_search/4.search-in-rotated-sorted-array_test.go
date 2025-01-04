package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name     string
		input    args
		expected int
	}{
		{"case1", args{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
		{"case2", args{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1},
		{"case3", args{[]int{1}, 0}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, search(tt.input.nums, tt.input.target))
		})
	}
}
