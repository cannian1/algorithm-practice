package skill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDuplicate(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{"case1", []int{1, 3, 4, 2, 2}, 2},
		{"case2", []int{3, 1, 3, 4, 2}, 3},
		{"case3", []int{3, 3, 3, 3, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findDuplicate(tt.args))
		})
	}
}
