package two_points

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{"case1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"case2", []int{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxArea(tt.args), "maxArea(%v)", tt.args)
		})
	}
}
