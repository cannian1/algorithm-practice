package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decodeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"case1", "3[a]2[bc]", "aaabcbc"},
		{"case2", "3[a2[c]]", "accaccacc"},
		{"case3", "2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"case4", "abc3[cd]xyz", "abccdcdcdxyz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, decodeString(tt.input))
		})
	}
}
