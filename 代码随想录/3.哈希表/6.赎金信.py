# 383. 赎金信
# https://leetcode.cn/problems/ransom-note/


class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        ransom_count = [0] * 26
        magazine_count = [0] * 26

        for c in ransomNote:
            ransom_count[ord(c) - ord('a')] += 1
        for c in magazine:
            magazine_count[ord(c) - ord('a')] += 1
        return all(ransom_count[i] <= magazine_count[i] for i in range(26))


class Solution2:
    """使用字典"""

    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        hashmap = dict()
        for val in magazine:
            hashmap[val] = hashmap.get(val, 0) + 1

        for val in ransomNote:
            if val in hashmap and hashmap[val] > 0:
                hashmap[val] -= 1
            else:
                return False
        return True


class Solution3:
    """使用 defaultdict """

    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        from collections import defaultdict
        hmap = defaultdict(lambda: 0)
        for val in magazine:
            hmap[val] += 1

        for val in ransomNote:
            if val not in hmap or hmap[val] == 0:
                return False
            hmap[val] -= 1
        return True
