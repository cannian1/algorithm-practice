package skill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			"case1",
			[]int{1, 2, 3, 2, 1},
			3,
		},
		{
			"case2",
			[]int{5, 4, 5, 4, 1},
			1,
		},
		{
			"case3",
			[]int{2, 3, 3, 6, 6},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, singleNumber(tt.args), "singleNumber(%v)", tt.args)
		})
	}
}
