package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_spiralOrder(t *testing.T) {

	tests := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name: "case1",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "case2",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, spiralOrder(tt.input), "spiralOrder(%v)", tt.input)
		})
	}
}
