package two_points

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSum(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want [][]int
	}{
		{"case1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"case2", []int{0, 1, 1}, [][]int{}},
		{"case1", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, threeSum(tt.args), "threeSum(%v)", tt.args)
		})
	}
}
