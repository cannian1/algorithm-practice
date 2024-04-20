# 3. 无重复字符的最长子串 https://leetcode.cn/problems/longest-substring-without-repeating-characters/

# 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
"""枚举右端点，收缩左端点"""


class Solution:
    """滑动窗口"""

    def lengthOfLongestSubstring(self, s: str) -> int:
        result = left = 0
        window = set()  # 维护从下标 left 到下标 right 的字符
        for right, c in enumerate(s):
            while c in window:  # 如果加入 c 会导致窗口内有重复元素，就把滑动窗口最左侧的元素去除
                window.remove(s[left])
                left += 1  # 缩小窗口
            window.add(c)  # 加入 c
            result = max(result, right - left + 1)  # 更新窗口最大值
        return result
