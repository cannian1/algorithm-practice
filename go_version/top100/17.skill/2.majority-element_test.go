package skill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{"case1", []int{3, 2, 3}, 3},
		{"case2", []int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, majorityElement(tt.args), "majorityElement(%v)", tt.args)
		})
	}
}
