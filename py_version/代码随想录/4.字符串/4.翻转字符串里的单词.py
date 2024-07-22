# 151. 反转字符串中的单词
# https://leetcode.cn/problems/reverse-words-in-a-string/
class Solution:
    def reverseWords(self, s: str) -> str:
        # 删除前后空白
        s = s.strip()
        # 反转整个字符串
        s = s[::-1]
        # 将字符串拆分为单词，并反转每个单词
        s = ' '.join(word[::-1] for word in s.split())
        return s


class Solution2:
    """用内置函数"""

    def reverseWords(self, s: str) -> str:
        return ' '.join(reversed(s.split()))


class Solution3:
    """双指针"""

    def reverseWords(self, s: str) -> str:
        # 将字符串拆分为单词，即转换成列表类型
        words = s.split()

        # 反转单词
        left, right = 0, len(words) - 1
        while left < right:
            words[left], words[right] = words[right], words[left]
            left += 1
            right -= 1

        # 将列表转换成字符串
        return " ".join(words)


# 迭代器是一个惰性序列，它只在需要时才产生值，下面这种做法是将其类型转换为字符串
a = "abcd"
reversed_a = ''.join(reversed(a))
print(reversed_a)
list_re_a = list(reversed(a))  # 也可以转化为其他类型
print(list_re_a)
