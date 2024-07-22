# https://leetcode.cn/problems/happy-number/
# 202. 快乐数

class Solution:
    """使用集合"""

    def isHappy(self, n: int) -> bool:
        record = set()

        while True:
            n = self.get_sum(n)
            if n == 1:
                return True
            # 中间结果如果曾经出现过，那么说明陷入死循环，这个数不是快乐数
            if n in record:
                return False
            else:
                record.add(n)

    def get_sum(self, n: int) -> int:
        new_num = 0
        while n:
            # r = n % 10
            # n = n //10
            n, r = divmod(n, 10)
            new_num += r ** 2
        return new_num


class Solution2:
    """
    使用集合
    Bad：整型转字符串，字符串再转整型运算
    """

    def isHappy(self, n: int) -> bool:
        record = set()
        while n not in record:
            record.add(n)
            new_num = 0

            n_str = str(n)
            for i in n_str:
                new_num += int(i) ** 2

            if new_num == 1:
                return True
            else:
                n = new_num
        return False


class Solution3:
    """快慢指针"""

    def isHappy(self, n: int) -> bool:
        slow, fast = n, n
        while self.get_sum(n) != 1 and self.get_sum(self.get_sum(fast)):
            slow = self.get_sum(slow)
            fast = self.get_sum(self.get_sum(fast))

            if slow == fast:
                return False
        return True

    def get_sum(self, n: int) -> int:
        new_num = 0
        while n:
            # r = n % 10
            # n = n //10
            n, r = divmod(n, 10)
            new_num += r ** 2
        return new_num
