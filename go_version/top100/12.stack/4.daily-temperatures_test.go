package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dailyTemperatures(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"case1", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"case2", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{"case3", []int{60, 60, 90}, []int{2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, dailyTemperatures(tt.input))
		})
	}
}
