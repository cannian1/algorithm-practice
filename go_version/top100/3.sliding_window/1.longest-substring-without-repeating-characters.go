// 3. 无重复字符的最长子串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

package sliding_window

// 给定一个字符串 s ,请你找出其中不含有重复字符的最长子串的长度

func lengthOfLongestSubstring(s string) int {
	result, left := 0, 0
	window := make(map[rune]bool) // 使用 map 来维护从下标 left 到下标 right 的字符

	for right, c := range s {
		for window[c] { // 如果加入 c 会导致窗口内有重复元素，就把滑动窗口最左侧的元素去除
			delete(window, rune(s[left])) // 从窗口中移除字符
			left++                        // 缩小窗口
		}
		window[c] = true                   // 加入 c 到窗口
		result = max(result, right-left+1) // 更新窗口最大值
	}
	return result
}

/*  如果想要返回具体的子串，把 result 那行改成这个
if right-left+1 > result {
	result = right - left + 1
	start = left // 更新最长子串的起始位置
}
*/
