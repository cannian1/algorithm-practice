# 1047. 删除字符串中的所有相邻重复项
# https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string

class Solution:
    def removeDuplicates(self, s: str) -> str:
        res = []
        for item in s:
            # 栈不为空且栈顶与遍历到的元素相同
            if res and res[-1] == item:
                res.pop()
            else:
                res.append(item)
        return ''.join(res)  # 字符串拼接
