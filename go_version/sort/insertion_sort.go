package sort

// InsertionSort 插入排序 O(n²)
func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		base := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = base
	}
}
