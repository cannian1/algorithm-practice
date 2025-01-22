package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_invertTree(t *testing.T) {

	tests := []struct {
		name     string
		input    *TreeNode
		expected *TreeNode
	}{
		{
			"case1",
			&TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			&TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			},
		},
		{
			"case2",
			&TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			&TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, invertTree(tt.input), "invertTree(%v)", tt.input)
		})
	}
}
