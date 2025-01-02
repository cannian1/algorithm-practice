package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name     string
		input    args
		expected int
	}{
		{"case1", args{[]int{3, 2, 1, 5, 6, 4}, 2}, 5},
		{"case2", args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, findKthLargest(tt.input.nums, tt.input.k))
		})
	}
}
