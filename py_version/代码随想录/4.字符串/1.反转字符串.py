# 344. 反转字符串
# https://leetcode.cn/problems/reverse-string/

from typing import List


class Solution:
    def reverseString(self, s: List[str]) -> None:
        n = len(s)
        mid = n // 2
        for i in range(mid):
            s[i], s[n - 1 - i] = s[n - 1 - i], s[i]


class Solution2:
    def reverseString(self, s: List[str]) -> None:
        left, right = 0, len(s) - 1
        while left < right:
            s[left], s[right] = s[right], s[left]
            left += 1
            right -= 1


class Solution3:
    """栈"""

    def reverseString(self, s: List[str]) -> None:
        stack = []
        for c in s:
            stack.append(c)
        for i in range(len(s)):
            s[i] = stack.pop()
