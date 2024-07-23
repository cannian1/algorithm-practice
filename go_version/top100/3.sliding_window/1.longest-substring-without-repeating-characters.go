// 3. 无重复字符的最长子串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

package sliding_window

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	window := [128]bool{}
	left := 0
	result := 0

	for right, c := range s {
		for window[c] {
			window[s[left]] = false
			left++
		}
		window[c] = true
		result = max(result, right-left+1)
	}
	return result
}
