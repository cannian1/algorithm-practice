package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name     string
		input    args
		expected []int
	}{
		{"case1", args{[]int{1, 1, 1, 2, 2, 3}, 2}, []int{1, 2}},
		{"case2", args{[]int{1}, 1}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, topKFrequent(tt.input.nums, tt.input.k))
		})
	}
}
