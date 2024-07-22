# 541. 反转字符串 II
# https://leetcode.cn/problems/reverse-string-ii/

class Solution:
    def reverseStr(self, s: str, k: int) -> str:
        def reverse_sub_string(sub_s: list):
            n = len(sub_s)
            mid = n // 2
            for i in range(mid):
                sub_s[i], sub_s[n - 1 - i] = sub_s[n - 1 - i], sub_s[i]
            return sub_s

        res = list(s)

        # 每 2k 的步长，都要对前 k 个元素反转
        for cur in range(0, len(s), 2 * k):
            res[cur:cur + k] = reverse_sub_string(res[cur:cur + k])

        return ''.join(res)


class Solution2:
    def reverseStr(self, s: str, k: int) -> str:
        # Two pointers. Another is inside the loop.
        p = 0
        while p < len(s):
            p2 = p + k
            # Written in this could be more pythonic.
            s = s[:p] + s[p: p2][::-1] + s[p2:]
            p = p + 2 * k
        return s
