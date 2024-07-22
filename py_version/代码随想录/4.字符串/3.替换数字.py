# 替换数字
# https://www.programmercarl.com/kama54.%E6%9B%BF%E6%8D%A2%E6%95%B0%E5%AD%97.html

class Solution:
    def change(self, s):
        lst = list(s)  # Python里面的string也是不可改的，所以也是需要额外空间的。空间复杂度：O(n)。
        for i in range(len(lst)):
            if lst[i].isdigit():
                lst[i] = "number"
        return ''.join(lst)
   