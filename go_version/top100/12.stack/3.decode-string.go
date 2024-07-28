// 394.字符串解码
// https://leetcode.cn/problems/decode-string/

package stack

import "unicode"

// 用于保存栈中每个重复次数和对应的字符串状态
type pair struct {
	multi int
	res   string
}

func decodeString(s string) string {
	// 保存需要重复的次数和之前的结果字符串
	var stack []pair
	res := ""
	multi := 0

	for _, c := range s {
		// 遇到 [ ，将当前 multi 压栈，并重置它们
		if c == '[' {
			stack = append(stack, pair{multi, res})
			res = ""
			multi = 0
		} else if c == ']' { // 遇到 ] ，从栈中弹出上一个状态，进行重复并拼接
			// 辅助栈弹出元素
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = last.res + repeatString(res, last.multi)
		} else if unicode.IsDigit(c) { // c 为数字字符时，转化为数字并更新 multi 用于后续计算
			multi = multi*10 + int(c-'0')
		} else { // 遇到字母时，追加到 res 中
			res += string(c)
		}
	}
	return res
}

func repeatString(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
