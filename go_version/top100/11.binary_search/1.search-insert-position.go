// 35.搜索插入位置
// https://leetcode.cn/problems/search-insert-position/

package binary_search

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func searchInsert2(nums []int, target int) int {
	return commonSearch(len(nums), func(i int) bool {
		return nums[i] >= target
	})
}
