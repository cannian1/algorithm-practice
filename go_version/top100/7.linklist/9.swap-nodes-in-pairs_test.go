package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_swapPairs(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"case1", []int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{"case2", []int{1, 2, 3}, []int{2, 1, 3}},
		{"case3", []int{}, []int{}},
		{"case4", []int{1}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, ListToSlice(swapPairs(SliceToList(tt.input))),
				"swapPairs(%v)", tt.input)
		})
	}
}
