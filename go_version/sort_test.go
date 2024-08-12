package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestInsertionSort(t *testing.T) {
	arr := []int{5, 7, 3, 1, 2}
	expectArr := []int{1, 2, 3, 5, 7}
	BubbleSort(arr)
	assert.Equal(t, expectArr, arr)
}

func TestMergeSort(t *testing.T) {
	arr := []int{5, 7, 3, 1, 2}
	expectArr := []int{1, 2, 3, 5, 7}
	res := MergeSort(arr)
	assert.Equal(t, expectArr, res)
}
