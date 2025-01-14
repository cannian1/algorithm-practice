package substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name     string
		input    args
		expected int
	}{
		{"case1", args{[]int{1, 1, 1}, 2}, 2},
		{"case2", args{[]int{1, 2, 3}, 3}, 2},
		{"case3", args{[]int{2, 4, 3}, 7}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, subarraySum(tt.input.nums, tt.input.k), "subarraySum(%v, %v)", tt.input.nums, tt.input.k)
		})
	}
}
