package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {

	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name: "case1",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "case2",
			input: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			rotate(tt.input)
			assert.Equal(t, tt.expected, input)
		})
	}
}
