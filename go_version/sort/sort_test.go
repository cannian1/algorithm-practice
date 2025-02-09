package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectSort(t *testing.T) {
	arr := []int{5, 7, 3, 1, 2}
	expectArr := []int{1, 2, 3, 5, 7}
	SelectSort(arr)
	assert.Equal(t, expectArr, arr)
}

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 7, 3, 1, 2}
	expectArr := []int{1, 2, 3, 5, 7}
	BubbleSort(arr)
	assert.Equal(t, expectArr, arr)
}

func TestQuickSort(t *testing.T) {
	arr := []int{5, 7, 3, 1, 2}
	expectArr := []int{1, 2, 3, 5, 7}
	res := QuickSort(arr)
	assert.Equal(t, expectArr, res)
}

func TestQuickSort_(t *testing.T) {
	arr := []int{5, 7, 3, 3, 1, 2}
	expectArr := []int{1, 2, 3, 3, 5, 7}
	res := QuickSort_(arr)
	assert.Equal(t, expectArr, res)
}

func TestQuickSort2(t *testing.T) {
	arr := []int{5, 7, 3, 3, 1, 2}
	expectArr := []int{1, 2, 3, 3, 5, 7}
	quickSort(arr, 0, len(arr)-1)
	assert.Equal(t, expectArr, arr)
}

func TestInsertionSort(t *testing.T) {
	arr := []int{5, 7, 3, 1, 2}
	expectArr := []int{1, 2, 3, 5, 7}
	InsertionSort(arr)
	assert.Equal(t, expectArr, arr)
}

func TestMergeSort(t *testing.T) {
	arr := []int{5, 7, 3, 1, 2}
	expectArr := []int{1, 2, 3, 5, 7}
	res := MergeSort(arr)
	assert.Equal(t, expectArr, res)
}
