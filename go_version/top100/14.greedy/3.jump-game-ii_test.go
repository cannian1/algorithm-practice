package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jump(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"case1", []int{2, 3, 1, 1, 4}, 2},
		{"case2", []int{2, 3, 0, 1, 4}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, jump(tt.input), "jump(%v)", tt.input)
		})
	}
}
