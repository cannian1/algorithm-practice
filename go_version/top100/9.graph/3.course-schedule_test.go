package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name     string
		input    args
		expected bool
	}{
		{
			"case1",
			args{
				numCourses: 4,
				prerequisites: [][]int{
					{1, 0},
					{2, 1},
					{3, 2},
				},
			},
			true,
		},
		{
			"case2",
			args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.expected, canFinish(tt.input.numCourses, tt.input.prerequisites),
				"canFinish(%v, %v)", tt.input.numCourses, tt.input.prerequisites)
		})
	}
}
