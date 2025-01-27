package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsets(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{"case1", []int{1, 2, 3}, [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}},
		{"case2", []int{0}, [][]int{{}, {0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, subsets(tt.input), "subsets(%v)", tt.input)
		})
	}
}
