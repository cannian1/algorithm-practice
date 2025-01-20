package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head []int
		k    int
	}
	tests := []struct {
		name     string
		input    args
		expected []int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 5}, 2}, []int{2, 1, 4, 3, 5}},
		{"case1", args{[]int{1, 2, 3, 4, 5}, 3}, []int{3, 2, 1, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, ListToSlice(reverseKGroup(SliceToList(tt.input.head), tt.input.k)),
				"reverseKGroup(%v, %v)", tt.input.head, tt.input.k)
		})
	}
}
