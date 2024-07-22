# 150. 逆波兰表达式求值
# https://leetcode.cn/problems/evaluate-reverse-polish-notation/

from typing import List
from operator import add, sub, mul


class Solution:
    op_map = {'+': add, '-': sub, '*': mul, '/': lambda x, y: int(x / y)}

    def evalRPN(self, tokens: List[str]) -> int:
        stack = []
        for token in tokens:
            if stack not in {'+', '-', '*', '/'}:
                stack.append(int(token))
            else:
                op2 = stack.pop()
                op1 = stack.pop()
                stack.append(self.op_map[token](op1, op2))  # 先出栈的在后面
        return stack.pop()
