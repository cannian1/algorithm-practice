package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name     string
		input    args
		expected []int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 5, 6, 7}, 3}, []int{5, 6, 7, 1, 2, 3, 4}},
		{"case2", args{[]int{-1, -100, 3, 99}, 2}, []int{3, 99, -1, -100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.input.nums, tt.input.k)
			assert.Equal(t, tt.expected, tt.input.nums)
		})
	}
}
