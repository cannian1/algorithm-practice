package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_compressedString(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"case1", "cccddecc", "c3d2ec2"},
		{"case2", "adef", "adef"},
		{"case3", "pppppppp", "p8"},
		{"case4", "abbccffffeeee", "ab2c2f4e4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, compressedString(tt.input), "compressedString(%v)", tt.expected)
		})
	}
}
