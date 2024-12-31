package sliding_window

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"case1", "abcabcbb", 3},
		{"case2", "bbbbb", 1},
		{"case3", "pwwkew", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, lengthOfLongestSubstring(tt.input))
		})
	}
}
