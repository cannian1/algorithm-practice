// 49. 字母异位词分组
// https://leetcode.cn/problems/group-anagrams/

package hash

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)

	for _, str := range strs {
		s := []byte(str)
		slices.Sort(s)
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}

	result := make([][]string, 0, len(mp))
	for _, v := range mp {
		result = append(result, v)
	}
	return result
}
