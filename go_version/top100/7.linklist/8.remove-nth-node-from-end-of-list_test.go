package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head []int
		n    int
	}
	tests := []struct {
		name     string
		input    args
		expected []int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 5}, 3}, []int{1, 2, 4, 5}},
		{"case2", args{[]int{1, 2}, 1}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, ListToSlice(removeNthFromEnd(SliceToList(tt.input.head), tt.input.n)),
				"removeNthFromEnd(%v, %v)", tt.input.head, tt.input.n)
		})
	}
}
