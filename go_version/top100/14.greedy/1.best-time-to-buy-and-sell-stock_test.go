package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"case1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"case2", []int{7, 6, 4, 3, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, maxProfit(tt.input), "maxProfit(%v)", tt.input)
		})
	}
}
