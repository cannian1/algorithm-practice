package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidBST(t *testing.T) {

	tests := []struct {
		name     string
		input    *TreeNode
		expected bool
	}{
		{
			"case1",
			&TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			}, true,
		},
		{
			"case2",
			&TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, isValidBST(tt.input), "isValidBST(%v)", tt.input)
		})
	}
}
