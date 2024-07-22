// 128. 最长连续序列
// https://leetcode.cn/problems/longest-consecutive-sequence/
package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums1 := []int{2, 7, 11, 15}
	nums2 := []int{3, 2, 4}
	nums3 := []int{3, 3}

	r1 := twoSum(nums1, 9)
	r2 := twoSum(nums2, 6)
	r3 := twoSum(nums3, 6)

	assert.Equal(t, r1, []int{0, 1})
	assert.Equal(t, r2, []int{1, 2})
	assert.Equal(t, r3, []int{0, 1})
}

func TestGroupAnagrams(t *testing.T) {
	arr1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	arr2 := []string{}
	arr3 := []string{"a"}

	r1 := groupAnagrams(arr1)
	r2 := groupAnagrams(arr2)
	r3 := groupAnagrams(arr3)

	t.Log(r1)
	t.Log(r2)
	t.Log(r3)
}

func TestLongestConsecutive(t *testing.T) {
	arr1 := []int{100, 4, 200, 1, 3, 2}
	arr2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}

	r1 := longestConsecutive(arr1)
	r2 := longestConsecutive(arr2)
	assert.Equal(t, r1, 4)
	assert.Equal(t, r2, 9)
}
