// 75.颜色分类
// https://leetcode.cn/problems/sort-colors

package skill

func sortColors(nums []int) {
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
	for i := 0; i < len(nums); i++ {
		if i < red {
			nums[i] = 0
		} else if i < red+white {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}
