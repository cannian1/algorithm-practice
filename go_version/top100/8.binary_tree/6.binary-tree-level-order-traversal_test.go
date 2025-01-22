package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_levelOrder(t *testing.T) {
	tests := []struct {
		name  string
		input *TreeNode
		want  [][]int
	}{
		{
			"case1",
			&TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			[][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, levelOrder(tt.input), "levelOrder(%v)", tt.input)
		})
	}
}
