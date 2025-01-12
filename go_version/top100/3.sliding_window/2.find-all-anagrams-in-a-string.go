// 438.找到字符串中所有字母异位词
// https://leetcode.cn/problems/find-all-anagrams-in-a-string/

package sliding_window

// 给定两个字符串 s 和 p，找到 s 中所有 p 的异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

func findAnagrams(s string, p string) []int {
	result := make([]int, 0)
	if len(s) < len(p) {
		return result
	}

	// 记录 p 中每个字符的频率
	pCount := [26]int{}
	for i := 0; i < len(p); i++ {
		pCount[p[i]-'a']++
	}

	// 用一个滑动窗口大小为 len(p) 来遍历 s
	sCount := [26]int{}
	for i := 0; i < len(p); i++ {
		sCount[s[i]-'a']++
	}

	// 比较初始窗口与 pCount 是否相同
	if sCount == pCount {
		result = append(result, 0)
	}

	// 滑动窗口
	for i := len(p); i < len(s); i++ {
		sCount[s[i]-'a']++        // 添加当前字符
		sCount[s[i-len(p)]-'a']-- // 移除窗口最左边的字符

		if sCount == pCount { // 比较窗口与 pCount 是否相同
			result = append(result, i-len(p)+1)
		}
	}

	return result
}
