package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSymmetric(t *testing.T) {
	tests := []struct {
		name  string
		input *TreeNode
		want  bool
	}{
		{"case1", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 3},
			},
		},
			true},
		{"case2", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
		},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isSymmetric(tt.input), "isSymmetric(%v)", tt.input)
		})
	}
}
