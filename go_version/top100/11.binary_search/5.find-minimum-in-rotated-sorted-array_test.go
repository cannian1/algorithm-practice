package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMin(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"case1", []int{3, 4, 5, 1, 2}, 1},
		{"case2", []int{4, 5, 6, 7, 0, 1, 2}, 0},
		{"case3", []int{11, 13, 15, 17}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, findMin(tt.input))
		})
	}
}
