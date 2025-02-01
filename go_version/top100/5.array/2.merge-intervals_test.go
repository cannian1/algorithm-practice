package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {

	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{"case1", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"case2", [][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{"case3", [][]int{{1, 8}, {4, 5}}, [][]int{{1, 8}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, merge(tt.input))
		})
	}
}
