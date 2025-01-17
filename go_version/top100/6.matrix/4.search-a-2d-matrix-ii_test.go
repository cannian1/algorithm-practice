package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name     string
		input    args
		expected bool
	}{
		{
			name: "case1",
			input: args{[][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			}, 5},
			expected: true,
		},
		{
			name: "case2",
			input: args{[][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			}, 20},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, searchMatrix(tt.input.matrix, tt.input.target), "searchMatrix(%v, %v)",
				tt.input.matrix, tt.input.target)
		})
	}
}
