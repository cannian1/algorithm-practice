# 20.有效的括号
# https://leetcode.cn/problems/valid-parentheses

class Solution:
    def isValid(self, s: str) -> bool:
        if len(s)% 2 != 0:
            return False
        stack = []
        for v in s:
            # 遇到左括号就把右括号入栈，匹配时直接弹出来就可以
            if v == '(':
                stack.append(')')
            elif v == '[':
                stack.append(']')
            elif v == '{':
                stack.append('}')
                # 空栈和栈顶元素与当前遍历的元素不匹配时返回失败
            elif not stack or stack[-1] != v:
                return False
            else:
                stack.pop()

        # 字符串遍历完了如果栈不为空，那就是不匹配
        return True if not stack else False