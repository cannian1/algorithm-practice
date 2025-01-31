package others

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name     string
		input    args
		expected string
	}{
		{"case1", args{"456", "77"}, "533"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, addStrings(tt.input.num1, tt.input.num2))
		})
	}
}
