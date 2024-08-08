// 438.找到字符串中所有字母异位词
// https://leetcode.cn/problems/find-all-anagrams-in-a-string/

package sliding_window

func findAnagrams(s string, p string) []int {
	need := make(map[byte]int)   // 记录 p 中所有字符出现次数
	window := make(map[byte]int) // 记录窗口中所有字符出现次数

	for i := range p {
		need[p[i]]++
	}

	left, right := 0, 0      // 当前窗口的左右边界
	valid := 0               // 记录窗口中满足 need 条件的字符个数
	result := make([]int, 0) // 用来存储所有找到的字母异位词的起始索引

	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) {
				result = append(result, left)
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return result
}
