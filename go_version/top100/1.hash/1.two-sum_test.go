package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{nums: []int{2, 7, 11, 15}, target: 9}, []int{0, 1}},
		{"case2", args{nums: []int{3, 2, 4}, target: 6}, []int{0, 1}},
		{"case3", args{nums: []int{3, 3}, target: 6}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, twoSum(tt.args.nums, tt.args.target), "twoSum(%v, %v)", tt.args.nums, tt.args.target)
		})
	}
}
