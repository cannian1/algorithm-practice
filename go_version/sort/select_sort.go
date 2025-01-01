package sort

// SelectSort 选择排序
func SelectSort(nums []int) {
	n := len(nums)
	for i := range n - 1 {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}
