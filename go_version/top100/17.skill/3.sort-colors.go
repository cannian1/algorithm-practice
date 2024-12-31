// 75.颜色分类
// https://leetcode.cn/problems/sort-colors

package skill

func sortColors(nums []int) {
	// 双指针法 O(n)、O(1)
	if len(nums) == 0 {
		return
	}

	left, right := 0, len(nums)-1 // 分别指向数组的开始和结束
	current := 0                  // 当前指针

	for current <= right {
		switch nums[current] {
		case 0:
			// 如果当前是0，将其与左指针交换，左指针和当前指针向右移动
			nums[current], nums[left] = nums[left], nums[current]
			left++
			current++
		case 1:
			// 如果当前是1，跳过
			current++
		case 2:
			// 如果当前是2，将其与右指针交换，右指针向左移动
			nums[current], nums[right] = nums[right], nums[current]
			right--
		}
	}
}

// 计数排序 O(n)、O(n)
func sortColors1(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}

	red, white, blue := 0, 0, 0
	for _, color := range nums {
		switch color {
		case 0:
			red++
		case 1:
			white++
		case 2:
			blue++
		}
	}
	for i := 0; i < n; i++ {
		if i < red {
			nums[i] = 0
		} else if i < red+white {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}
