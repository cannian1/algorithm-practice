# 经典排序
def selection_sort(nums: list[int]):
    """选择排序 O(n²) O(1) 不稳定"""
    n = len(nums)

    # 每轮选出待排序区间最小值与待排序区间第一个索引处的元素交换
    for i in range(n - 1):
        min_val = i
        for j in range(i + 1, n):
            if nums[j] < nums[min_val]:
                min_val = j
        nums[i], nums[min_val] = nums[min_val], nums[i]


def bubble_sort(nums: list[int]):
    """冒泡排序 O(n²) O(1) 稳定"""
    n = len(nums)

    # 每轮把最大的泡泡冒到水面
    for i in range(n - 1, 0, -1):
        flag = False
        for j in range(i):
            if nums[j] > nums[j + 1]:
                nums[j], nums[j + 1] = nums[j + 1], nums[i]
                flag = True
        if not flag:  # 此轮冒泡未交换元素，跳出循环
            break


def insertion_sort(nums: list[int]):
    """插入排序 O(n²) O(1) 稳定"""

    # 初始状态默认第一个元素有序
    # 每次 i 向后遍历相当于向已排序的扑克牌插入新的扑克
    for i in range(1, len(nums)):
        base = nums[i]
        j = i - 1
        while j >= 0 and nums[j] > base:
            nums[j + 1] = nums[j]
            j -= 1
        nums[j + 1] = base


class QuickSort:
    """快速排序 O(nlogn) O(n) 不稳定"""

    def partition(self, nums: list[int], left: int, right: int) -> int:
        """哨兵划分"""
        # 以 nums[left] 为基数
        i, j = left, right
        while i < j:
            while i < j and nums[j] >= nums[left]:
                j -= 1
            while i < j and nums[i] < nums[left]:
                i += 1
            # 元素交换
            nums[i], nums[j] = nums[j], nums[i]
        # 将基准数交换至两子数组的分界线
        nums[i], nums[left] = nums[left], nums[i]
        return i

    def quick_sort(self, nums: list[int], left: int, right: int):
        """快速排序"""
        # 子数组长度为 1 时终止递归
        if left >= right:
            return
        # 哨兵划分
        pivot = self.partition(nums, left, right)
        # 递归左子数组、右子数组
        self.quick_sort(nums, left, pivot - 1)
        self.quick_sort(nums, pivot + 1, right)


def quicksort(nums: list[int]) -> list[int]:
    """快速排序，简单写法"""
    if len(nums) < 2:
        return nums

    pivot = nums[0]
    left, right = [], []

    for val in nums[1:]:
        if val <= pivot:
            left.append(val)
        else:
            right.append(val)

    return quicksort(left) + [pivot] + quicksort(right)


s = QuickSort()
arr = [4, 7, 9, 12, 7, 3, 1, 8]
s.quick_sort(arr, 0, len(arr) - 1)
print(arr)
