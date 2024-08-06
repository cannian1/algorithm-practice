def find_first_occurrence(nums: list[int], target: int) -> int:
    """查找非严格递增数组的第一次出现的目标元素的索引 O(logN)"""
    left, right = 0, len(nums) - 1
    result = -1
    while left <= right:
        mid = (left + right) // 2
        if nums[mid] > target:
            right = mid - 1
        elif nums[mid] < target:
            left = mid + 1
        else:
            # 当 array[mid] == target 时
            result = mid
            right = mid - 1  # 继续查找左侧部分
    return result


# 示例使用
arr = [1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 5, 6, 8, 8, 8]
index = find_first_occurrence(arr, 2)
print(index)  # 输出应该是 3
