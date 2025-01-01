// 20.有效的括号
// https://leetcode.cn/problems/valid-parentheses

package stack

// isValid 检查输入字符串的括号是否有效
func isValid(s string) bool {
	// 创建一个栈用于保存左括号
	var stack []rune

	// 定义一个映射，用于匹配右括号与对应的左括号
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 遍历字符串中的每个字符
	for _, c := range s {
		// 如果是右括号
		if leftBracket, exists := bracketMap[c]; exists {
			// 检查栈是否非空且栈顶是否匹配对应的左括号
			if len(stack) > 0 && stack[len(stack)-1] == leftBracket {
				// 从栈中弹出栈顶元素
				stack = stack[:len(stack)-1]
			} else {
				// 如果不匹配，则字符串无效
				return false
			}
		} else {
			// 如果是左括号，将其压入栈中
			stack = append(stack, c)
		}
	}
	// 如果栈为空，说明所有括号都匹配；否则字符串无效
	return len(stack) == 0
}
