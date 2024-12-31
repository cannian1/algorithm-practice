package two_points

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{"case1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"case2", []int{0}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args)
			assert.Equal(t, tt.want, tt.args)
		})
	}
}
