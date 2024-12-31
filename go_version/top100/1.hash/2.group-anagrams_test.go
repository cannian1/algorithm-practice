package hash

import (
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
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
