package substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name     string
		input    args
		expected []int
	}{
		{"case1", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3}, []int{3, 3, 5, 5, 6, 7}},
		{"case2", args{[]int{1}, 1}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, maxSlidingWindow(tt.input.nums, tt.input.k), "maxSlidingWindow(%v, %v)", tt.input.nums, tt.input.k)
		})
	}
}
