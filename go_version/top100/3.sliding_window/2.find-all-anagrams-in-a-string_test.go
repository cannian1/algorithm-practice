package sliding_window

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name     string
		input    args
		expected []int
	}{
		{"case1", args{"cbaebabacd", "abc"}, []int{0, 6}},
		{"case2", args{"abab", "ab"}, []int{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, findAnagrams(tt.input.s, tt.input.p), "findAnagrams(%v, %v)", tt.input.s, tt.input.p)
		})
	}
}
