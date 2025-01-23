package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rightSideView(t *testing.T) {

	tests := []struct {
		name     string
		input    *TreeNode
		expected []int
	}{
		{
			"case1",
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
			},
			[]int{1, 3, 4},
		},
		{
			"case2",
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}},
				},
				Right: &TreeNode{Val: 3},
			},
			[]int{1, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, rightSideView(tt.input), "rightSideView(%v)", tt.input)
		})
	}
}
