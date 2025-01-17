package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setZeroes(t *testing.T) {

	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{name: "case1", input: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{name: "case2", input: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := tt.input
			setZeroes(input)
			assert.Equal(t, tt.expected, input)
		})
	}
}
