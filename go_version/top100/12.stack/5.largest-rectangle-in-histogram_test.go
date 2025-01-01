package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestRectangleArea(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"case1", []int{2, 1, 5, 6, 2, 3}, 10},
		{"case2", []int{2, 4}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, largestRectangleArea(tt.input))
		})
	}
}
