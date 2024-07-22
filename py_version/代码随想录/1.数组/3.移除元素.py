# leetcode 27 移除元素
# https://leetcode.cn/problems/remove-element/

class Solution:
    def removeElement(self, nums: list[int], val: int) -> int:
        """
        :type nums: List[int]
        :type val: int
        :rtype: int
        """
        k: int = 0  # 慢指针
        # i 相当于快指针，快指针遍历到的元素赋给慢指针索引的位置，遇到要删除的元素就跳过
        for i, v in enumerate(nums):
            if v != val:
                nums[k] = v
                k += 1
        return k


nums = [3, 2, 2, 3]
e = Solution()
r = e.removeElement(nums, 3)
print(r)
print(nums)