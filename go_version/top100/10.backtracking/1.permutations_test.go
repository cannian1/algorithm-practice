package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_permute(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{"case1", []int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{"case2", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, permute(tt.input), "permute(%v)", tt.input)
		})
	}
}
