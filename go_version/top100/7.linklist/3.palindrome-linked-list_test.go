package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"case1", []int{1, 2, 3, 2, 1}, true},
		{"case2", []int{1, 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, isPalindrome(SliceToList(tt.input)), "isPalindrome(%v)", tt.input)
		})
	}
}
