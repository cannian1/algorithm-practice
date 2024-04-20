# 438.找到字符串中所有字母异位词 https://leetcode.cn/problems/find-all-anagrams-in-a-string
from collections import deque
from typing import List


class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        if len(s) < len(p):
            return []

        # 创建数组保存字符串 p 内所存各个元素的个数
        count_p = [0] * 26
        for c in p:
            count_p[ord(c) - ord('a')] += 1
        # 创建滑动窗口
        window = deque([])
        # 创建用于和 p 比较的数组
        cur = [0] * 26
        result = []

        for i, c in enumerate(s):
            idx = ord(c) - ord('a')
            cur[idx] += 1 # 在刚好满足条件的时候，这里的 cur 列表和 count_p 列表就相等了，往 result 里添加完数据不用管滑动窗口就知道了索引
            if cur == count_p:
                result.append(i - len(p) + 1)  # 收集符合条件的索引
            # 向滑动窗口添加元素
            window.append(c)
            # 如果窗口的长度等于 p 的长度，或者当前字符 c 的计数超过了 p 中该字符的计数，需要从窗口的左侧移除一个字符。
            if window and len(window) == len(p) or cur[idx] > count_p[idx]:
                left = window.popleft()
                cur[ord(left) - ord('a')] -= 1
        return result


s = Solution()
r = s.findAnagrams("abcaabc", "abc")
print(r)
