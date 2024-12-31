package two_points

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trap(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{"case1", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"case2", []int{4, 2, 0, 3, 2, 5}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, trap(tt.args), "trap(%v)", tt.args)
		})
	}
}
