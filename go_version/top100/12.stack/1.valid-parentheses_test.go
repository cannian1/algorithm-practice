package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"case1", "()", true},
		{"case2", "()[]{}", true},
		{"case3", "(]", false},
		{"case4", "([)]", false},
		{"case5", "{[]}", true},
		{"casa6", "[0L", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isValid(tt.input))
		})
	}
}
