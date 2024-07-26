// 153.寻找旋转排序数组中的最小值
// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array

package binary_search

//func _findMin(nums []int) int {
//	left, right := 0, len(nums)-1
//	for left <= right {
//		// 中间的数和最右边的数比较，如果中间数 < 最右边的数，那么最小值在当前左半区间，
//		//  如果中间数 > 最右边的数，那么最小值在当前右半区间
//
//		mid := left + ((right - left) >> 1)
//
//		if nums[mid] < nums[len(nums)-1] {
//			right = mid - 1
//		} else {
//			left = mid + 1
//		}
//	}
//	return nums[left]
//}

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
