# 76. 最小覆盖子串 https://leetcode.cn/problems/minimum-window-substring
import collections


class Solution:
    def minWindow(self, s: str, t: str) -> str:
        from collections import defaultdict

        # 初始化需要的字符及其数量的字典
        need = defaultdict(int)
        for c in t:
            need[c] += 1

        # 初始化窗口中满足need条件的字符及其数量的字典
        window = defaultdict(int)
        # 表示窗口中满足need条件的字符个数
        valid = 0

        # 初始化窗口的左右指针
        left = right = 0

        # 记录最小覆盖子串的起始索引及长度
        start, length = 0, float('inf')

        while right < len(s):
            # c是将移入窗口的字符
            c = s[right]
            # 扩大窗口
            right += 1

            # 进行窗口内数据的一系列更新
            if c in need:
                window[c] += 1
                if window[c] == need[c]:
                    valid += 1

            # 判断左侧窗口是否要收缩
            while valid == len(need):
                # 更新最小覆盖子串
                if right - left < length:
                    start = left
                    length = right - left

                # d是将移出窗口的字符
                d = s[left]
                # 缩小窗口
                left += 1

                # 进行窗口内数据的一系列更新
                if d in need:
                    if window[d] == need[d]:
                        valid -= 1
                    window[d] -= 1

        # 返回最小覆盖子串
        return "" if length == float('inf') else s[start:start + length]



s = Solution()

s.minWindow("sdadsadasd", "ads")
