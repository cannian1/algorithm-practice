package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name     string
		input    args
		expected [][]int
	}{
		{"case1", args{[]int{2, 3, 6, 7}, 7}, [][]int{{2, 2, 3}, {7}}},
		{"case2", args{[]int{2, 3, 5}, 8}, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{"case3", args{[]int{2}, 1}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, combinationSum(tt.input.candidates, tt.input.target), "combinationSum(%v, %v)", tt.input.candidates, tt.input.target)
		})
	}
}
