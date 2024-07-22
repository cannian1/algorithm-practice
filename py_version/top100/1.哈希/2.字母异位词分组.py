# 49. 字母异位词分组
# https://leetcode.cn/problems/group-anagrams/
import collections
from typing import List


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        # defaultdict 在 key 不存在时不会抛出异常
        mp = collections.defaultdict(list)

        for st in strs:
            key = "".join(sorted(st)) # 将字母异位词排序后的字符串作为哈希表的 key
            mp[key].append(st)

        return list(mp.values())
