# 20. 有效的括号
# https://leetcode.cn/problems/valid-parentheses/

"""数量、类型、顺序要匹配"""


class Solution:
    """只用栈"""

    def isValid(self, s: str) -> bool:
        if len(s) % 2 != 0:
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


class Solution2:
    """使用字典"""

    def isValid(self, s: str) -> bool:
        if len(s) % 2 != 0:
            return False

        stack = []
        mapping = {
            '(': ')',
            '[': ']',
            '{': '}'
        }

        for item in s:
            if item in mapping.keys():
                stack.append(mapping[item])
            elif not stack or stack[-1] != item:
                return False
            else:
                stack.pop()
        return True if not stack else False
