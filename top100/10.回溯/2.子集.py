# 78.子集
# https://leetcode.cn/problems/subsets/
# 子集类回溯

class Solution:
    def subsets(self, nums):
        result = []
        path = []
        self.backtracking(nums, 0, path, result)
        return result

    def backtracking(self, nums, startIndex, path, result):
        result.append(path[:])  # 收集子集，要放在终止添加的上面，否则会漏掉自己
        if startIndex >= len(nums):  # 终止条件可以不加
            return
        for i in range(startIndex, len(nums)):
            path.append(nums[i])
            self.backtracking(nums, i + 1, path, result)
            path.pop()


class Solution2:
    """
    站在输入的角度思考
    每个数可以在子集中（选）
    也可以不在子集中（不选）
    叶子是答案
    """

    # https://www.bilibili.com/video/BV1mG4y1A7Gu
    def subsets(self, nums):
        result = []
        path = []
        n = len(nums)

        def dfs(i: int) -> None:
            if i == n:
                result.append(path.copy())
                return

            dfs(i + 1)

            path.append(nums[i])
            dfs(i + 1)
            path.pop()  # 恢复现场

        dfs(0)
        return result


class Solution3:
    """
    站在答案的角度思考
    枚举第一个数选谁……
    枚举第二个数选谁……
    每个节点都是答案

    注意：[1,2]和[2,1]是重复的子集
    为了避免重复，下一个数应大于当前选择的数
    """

    def subsets(self, nums):
        result = []
        path = []
        n = len(nums)

        def dfs(i: int) -> None:
            # 递归到的每个节点都是答案
            result.append(path.copy())
            # if i == n:
            #     return

            # 枚举当前要填的数字
            for j in range(i, n):
                path.append(nums[j])
                dfs(j + 1)
                path.pop()  # 恢复现场

        dfs(0)
        return result


nums = [1, 2, 3]
s = Solution3()
r = s.subsets(nums)
print(r)
