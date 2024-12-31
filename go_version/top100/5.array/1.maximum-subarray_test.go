package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"case1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"case2", []int{1}, 1},
		{"case3", []int{5, 4, -1, 7, 8}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, maxSubArray(tt.input))
		})
	}
}
