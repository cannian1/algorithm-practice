package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	arr1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	arr2 := []int{1}
	arr3 := []int{5, 4, -1, 7, 8}

	r1 := maxSubArray(arr1)
	r2 := maxSubArray(arr2)
	r3 := maxSubArray(arr3)

	assert.Equal(t, r1, 6)
	assert.Equal(t, r2, 1)
	assert.Equal(t, r3, 23)
}

func TestMerge(t *testing.T) {
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals2 := [][]int{{1, 4}, {4, 5}}
	r1 := merge(intervals1)
	r2 := merge(intervals2)

	assert.Equal(t, r1, [][]int{{1, 6}, {8, 10}, {15, 18}})
	assert.Equal(t, r2, [][]int{{1, 5}})
}

func TestRotate(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	arr2 := []int{-1, -100, 3, 99}

	rotate(arr1, 3)
	rotate(arr2, 2)

	assert.Equal(t, arr1, []int{5, 6, 7, 1, 2, 3, 4})
	assert.Equal(t, arr2, []int{3, 99, -1, -100})
}
