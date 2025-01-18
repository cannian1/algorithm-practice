package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseList(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"case1", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"case2", []int{1, 2}, []int{2, 1}},
		{"case3", []int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, ListToSlice(reverseList(SliceToList(tt.input))), "reverseList(%v)", tt.input)
		})
	}
}
