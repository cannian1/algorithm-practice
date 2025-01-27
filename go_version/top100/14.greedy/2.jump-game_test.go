package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canJump(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"case1", []int{2, 3, 1, 1, 4}, true},
		{"case2", []int{3, 2, 1, 0, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, canJump(tt.input), "canJump(%v)", tt.input)
		})
	}
}
