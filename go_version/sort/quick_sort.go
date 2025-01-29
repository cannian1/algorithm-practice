package sort

import (
	"math/rand/v2"
	"slices"
)

// QuickSort 快速排序
func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	randomIndex := rand.IntN(len(nums))
	nums[0], nums[randomIndex] = nums[randomIndex], nums[0]

	pivot := nums[0]
	var left, mid, right []int

	for _, val := range nums {
		if val < pivot {
			left = append(left, val)
		} else if val > pivot {
			right = append(right, val)
		} else {
			mid = append(mid, val)
		}
	}
	return slices.Concat(QuickSort(left), mid, QuickSort(right))
}

func medianOfThree(nums []int) int {
	a, b, c := nums[0], nums[len(nums)/2], nums[len(nums)-1]
	if (a < b && b < c) || (c < b && b < a) {
		return b
	} else if (b < a && a < c) || (c < a && a < b) {
		return a
	}
	return c
}

// 三分取中法优化的快速排序
func QuickSort_(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	// 三分取中法选择基准值
	pivot := medianOfThree(nums)

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

func quickSort(nums []int, left, right int) {
	if left < right {
		pivotIndex := partition(nums, left, right) // 获取基准位置
		quickSort(nums, left, pivotIndex-1)        // 递归排序左半部分
		quickSort(nums, pivotIndex+1, right)       // 递归排序右半部分
	}
}

// 优化后的分区函数（随机选择基准）
func partition(nums []int, left, right int) int {
	// 随机选择基准并交换到末尾
	randomIndex := rand.IntN(right-left+1) + left
	nums[randomIndex], nums[right] = nums[right], nums[randomIndex]
	pivot := nums[right]

	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
