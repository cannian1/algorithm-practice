package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestConsecutive(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{"case1", []int{100, 4, 200, 1, 3, 2}, 4},
		{"case2", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestConsecutive(tt.args), "longestConsecutive(%v)", tt.args)
		})
	}
}
