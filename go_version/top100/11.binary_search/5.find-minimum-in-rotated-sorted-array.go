// 153.寻找旋转排序数组中的最小值
// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array

package binary_search

func findMin(nums []int) int {
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

	return nums[left]
}
