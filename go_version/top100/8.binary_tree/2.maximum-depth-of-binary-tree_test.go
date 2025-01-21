package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		name     string
		input    *TreeNode
		expected int
	}{
		{
			name: "case1",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: 3,
		},
		{
			name: "case2",
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, maxDepth(tt.input), "maxDepth(%v)", tt.input)
		})
	}
}
