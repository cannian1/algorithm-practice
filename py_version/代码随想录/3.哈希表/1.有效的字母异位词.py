# 242. 有效的字母异位词
# https://leetcode.cn/problems/valid-anagram/

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        record = [0] * 26
        for i in s:
            record[ord(i) - ord("a")] += 1
        for i in t:
            record[ord(i) - ord("a")] -= 1
        for i in range(26):
            if record[i] != 0:
                return False
        return True


class Solution2:
    def isAnagram(self, s: str, t: str) -> bool:
        # defaultdict 访问不存在的 key 不会抛出异常
        from collections import defaultdict

        s_dict = defaultdict(int)
        t_dict = defaultdict(int)
        for val in s:
            s_dict[val] += 1

        for val in t:
            t_dict[val] += 1
        return s_dict == t_dict

    def dictMethod(self, s: str, t: str):
        # dict 的 get方法第二个参数是设置key不存在时的默认值，不设置则为 None
        s_dict = dict()
        t_dict = dict()
        for i in s:
            s_dict[i] = s_dict.get(i, 0) + 1
        for j in t:
            t_dict[j] = t_dict.get(j, 0) + 1


class Solution3:
    def isAnagram(self, s: str, t: str) -> bool:
        from collections import Counter
        a_count = Counter(s)
        b_count = Counter(t)
        return a_count == b_count


test = Solution()
res = test.isAnagram("abcc", "ccba")
print(res)  # True
