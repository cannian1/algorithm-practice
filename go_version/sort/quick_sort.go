package sort

import "slices"

// QuickSort 快速排序
func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]
	var left, right []int

	for _, val := range nums[1:] {
		if val <= pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}
	return slices.Concat(QuickSort(left), []int{pivot}, QuickSort(right))
}

func medianOfThree(a, b, c int) int {
	if a > b {
		if b > c {
			return b
		}
		if a > c {
			return c
		}
		return a
	}
	if a > c {
		return a
	}
	if b > c {
		return c
	}
	return b
}

// 三分取中法优化的快速排序
func QuickSort_(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	// 三分取中法选择基准值
	mid := len(nums) / 2
	pivot := medianOfThree(nums[0], nums[mid], nums[len(nums)-1])

	var left, right, equal []int
	for _, val := range nums {
		if val < pivot {
			left = append(left, val)
		} else if val > pivot {
			right = append(right, val)
		} else {
			equal = append(equal, val)
		}
	}

	// 合并时确保将等于基准值的元素添加一次
	return slices.Concat(QuickSort_(left), equal, QuickSort_(right))
}
