package two_points

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	arr1 := []int{0, 1, 0, 3, 12}
	arr2 := []int{0}
	moveZeroes(arr1)

	assert.Equal(t, arr1, []int{1, 3, 12, 0, 0})
	assert.Equal(t, arr2, []int{0})
}

func TestContainerWithMostWater(t *testing.T) {
	arr1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	arr2 := []int{1, 1}

	r1 := maxArea(arr1)
	r2 := maxArea(arr2)
	assert.Equal(t, r1, 49)
	assert.Equal(t, r2, 1)
}

func TestThreeSum(t *testing.T) {
	arr1 := []int{-1, 0, 1, 2, -1, -4}
	arr2 := []int{0, 1, 1}
	arr3 := []int{0, 0, 0}

	r1 := threeSum(arr1)
	r2 := threeSum(arr2)
	r3 := threeSum(arr3)

	assert.Equal(t, r1, [][]int{{-1, -1, 2}, {-1, 0, 1}})
	assert.Equal(t, r2, [][]int{})
	assert.Equal(t, r3, [][]int{{0, 0, 0}})
}
