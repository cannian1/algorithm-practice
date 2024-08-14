// 20.有效的括号
// https://leetcode.cn/problems/valid-parentheses

package stack

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	mp := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	st := make([]rune, 0)
	for _, c := range s {
		if v := mp[c]; v > 0 { // c 是左括号
			st = append(st, v) // 右括号入栈
		} else {
			if len(st) == 0 || st[len(st)-1] != c {
				return false // 没有左括号或左括号类型不匹配
			}
			st = st[:len(st)-1] // 出栈
		}
	}
	return len(st) == 0
}
