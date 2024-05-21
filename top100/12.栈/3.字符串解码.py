# 394.字符串解码
# https://leetcode.cn/problems/decode-string/
class Solution:
    def decodeString(self, s: str) -> str:
        stack = []  # 辅助栈
        res = ""
        multi = 0

        for c in s:
            if c == '[':
                stack.append([multi, res])
                res, multi = "", 0
            elif c == ']':
                cul_multi, last_res = stack.pop()
                res = last_res + cul_multi * res
            elif '0' <= c <= '9':  # c为数字字符时，转化为数字用于后续计算
                multi = multi * 10 + int(c)
            else:  # c为字母时，在 res 尾部添加 c
                res += c
        return res
