package sort

// BubbleSort 冒泡排序
func BubbleSort(nums []int) {
	n := len(nums)
	var flag bool
	for i := n - 1; i > 0; i-- {
		flag = false
		for j := range i {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
