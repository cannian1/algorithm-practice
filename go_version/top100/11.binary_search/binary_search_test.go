package binary_search

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	//nums1 := []int{3, 4, 5, 1, 2}
	//nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	//nums3 := []int{11, 13, 15, 17}
	//r1 := _findMin(nums1)
	//r2 := _findMin(nums2)
	//r3 := _findMin(nums3)
	//
	//assert.Equal(t, r1, 1)
	//assert.Equal(t, r2, 0)
	//assert.Equal(t, r3, 11)

	nums1 := []int{3, 1}
	r := findMin(nums1)
	fmt.Println(r)
}

func TestSearch(t *testing.T) {
	//nums1 := []int{4, 5, 6, 7, 0, 1, 2}
	//r1 := search(nums1, 0)

	nums2 := []int{3, 1}
	r2 := search(nums2, 1)
	//fmt.Println(r1)
	fmt.Println(r2)
}
