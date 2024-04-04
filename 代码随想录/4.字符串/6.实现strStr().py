# 28. 实现 strStr()
# https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/

# https://blog.csdn.net/v_july_v/article/details/7041827 KMP详解

""" KMP 算法
文本串 aabaabaaf
模式串 aabaaf
前缀：包含首字母，不包含尾字母的所有子串
后缀：包含尾字母，不包含首字母的所有子串
最长相等前后缀
a       0
aa      1
aab     0
aaba    1
aabaa   2
aabaaf  0
右边的 010120 就是模式串 aabaaf 的前缀表

aabaaf
    -
010120

在逐个比较是否匹配时发现不匹配了，就去找前面那个子串的最长相等前后缀
在这个后缀的后面不匹配了，就要去与其相等的前缀后面继续匹配（下标为最长相等前后缀的长度）
"""

"""
BM 算法 和 Sunday 算法的最坏时间复杂度比 KMP 算法高
Go 语言标准库使用的是 Rabin-Karp 算法
"""

class Solution:
    """前缀表（减一）
    减一的前缀表就是把前缀表逐项减一，不匹配时跳回索引位置
    """

    def strStr(self, haystack: str, needle: str) -> int:
        if not needle:
            return 0
        next = [0] * len(needle)
        self.getNext(next, needle)  # 初始化前缀表

        j = -1
        for i in range(len(haystack)):
            while j >= 0 and haystack[i] != needle[j + 1]:
                j = next[j]
            if haystack[i] == needle[j + 1]:
                j += 1
            if j == len(needle) - 1:
                return i - len(needle) + 1
        return -1

    # 获取前缀表
    def getNext(self, next: list[int], s):
        j = -1  # j 的起始位置比 i 前一位
        next[0] = j
        for i in range(1, len(s)):
            # j 与 i 不匹配时
            while j >= 0 and s[i] != s[j + 1]:
                j = next[j]
            if s[i] == s[j + 1]:
                j += 1
            next[i] = j


class Solution3(object):
    """ 暴力法 """

    def strStr(self, haystack, needle):
        """
        :type haystack: str
        :type needle: str
        :rtype: int
        """
        m, n = len(haystack), len(needle)
        for i in range(m):
            if haystack[i:i + n] == needle:
                return i
        return -1
