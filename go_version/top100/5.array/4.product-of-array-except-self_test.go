package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_productExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"case1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"case2", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, productExceptSelf(tt.input), "productExceptSelf(%v)", tt.input)
		})
	}
}
