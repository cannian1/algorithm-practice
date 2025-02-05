package binary_search

// 二分查找第一次出现的元素
func binarySearchFirst(nums []int, target int) int {
	idx := commonSearch(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if idx >= 0 && nums[idx] == target {
		return idx
	} else {
		return -1
	}
}

// 从标准库抄来改的
func commonSearch(n int, fn func(int) bool) int {
	left, right := 0, n-1
	for left <= right {
		mid := int(uint(left+right) >> 1)
		if !fn(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
