package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_diameterOfBinaryTree(t *testing.T) {

	tests := []struct {
		name     string
		input    *TreeNode
		expected int
	}{
		{"case1", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{Val: 3},
		}, 3},
		{"case2", &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 2},
		}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, diameterOfBinaryTree(tt.input), "diameterOfBinaryTree(%v)", tt.input)
		})
	}
}
