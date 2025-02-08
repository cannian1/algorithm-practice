// 3163.压缩字符串
// https://leetcode.cn/problems/string-compression-iii/

package history

import "strconv"

func compressedString(word string) string {
	// 1. 用一个变量，记录上次出现的元素
	// 2. 当元素再次出现，就计数器自增
	// 3. 在遍历到的元素与上次出现的历史元素不相同时，新的元素和数字追加到结果后面，重置计数器和最近元素

	if len(word) == 0 {
		return ""
	}

	lastChar := word[0]
	currentCount := 1
	res := ""

	for i := 1; i < len(word); i++ {
		if word[i] == lastChar {
			currentCount++
		} else {
			res += string(lastChar)
			if currentCount > 1 {
				res += strconv.Itoa(currentCount)
			}
			lastChar = word[i]
			currentCount = 1
		}
	}

	// 4. 补上最后一轮的元素
	res += string(lastChar)
	if currentCount > 1 {
		res += strconv.Itoa(currentCount)
	}

	return res
}
