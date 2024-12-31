package skill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextPermutation(t *testing.T) {
	tests := []struct {
		name   string
		args   []int
		expect []int
	}{
		{"case1", []int{1, 2, 3}, []int{1, 3, 2}},
		{"case2", []int{3, 2, 1}, []int{1, 2, 3}},
		{"case3", []int{1, 1, 5, 4, 3}, []int{1, 3, 1, 4, 5}},
		{"case4", []int{1, 2}, []int{2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args)
			assert.Equal(t, tt.expect, tt.args)
		})
	}
}
