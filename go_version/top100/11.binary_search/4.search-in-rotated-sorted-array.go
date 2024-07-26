// 33.搜索旋转排序数组
// https://leetcode.cn/problems/search-in-rotated-sorted-array/

package binary_search

func search(nums []int, target int) int {
	minIdx := findMinIdx(nums)

	// 最小值作为基准
	left, right := minIdx, minIdx+len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		i := mid % len(nums) // 取余映射
		if nums[i] < target {
			left = mid + 1
		} else if nums[i] > target {
			right = mid - 1
		} else {
			return i
		}
	}
	return -1
}

// 二分查找旋转排序数组最小值的索引
func findMinIdx(nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// 最小值在右半区间
			left = mid + 1
		} else if nums[mid] < nums[right] {
			// 最小值在左半区间
			right = mid
		} else {
			// 无法确定最小值在哪个半区间，右边界收缩
			right--
		}
	}
	return left
}
