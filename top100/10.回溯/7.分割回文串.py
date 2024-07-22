# 131.分割回文串
# https://leetcode.cn/problems/palindrome-partitioning
from typing import List


class Solution:
    """
    枚举逗号的位置，选/不选
    用子集型问题的思想
    """

    def partition(self, s: str) -> List[List[str]]:
        result = []
        path = []
        n = len(s)

        def dfs(i: int) -> None:
            if i == n:
                result.append(path.copy())
                return

            for j in range(i, n):  # 枚举子串的结束位置
                t = s[i:j + 1]
                if t == t[::-1]:  # 判断是否回文
                    path.append(t)
                    dfs(j + 1)
                    path.pop()  # 恢复现场

        dfs(0)
        return result


s = Solution()
r = s.partition("aab")
print(r)
