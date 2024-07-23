// 42.接雨水
// https://leetcode.cn/problems/trapping-rain-water/

package two_points

func trap(height []int) int {
	// 从柱子两侧开始向中间遍历
	result := 0
	left := 0
	right := len(height) - 1

	leftMax, rightMax := 0, 0

	for left < right {
		// 记录两侧最高高度
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		// 横坐标上某点的积水深度取决于左右指针最大值中较小的那个与当前该点的高度的差值
		if height[left] < height[right] {
			result += leftMax - height[left]
			left++
		} else {
			result += rightMax - height[right]
			right--
		}
	}
	return result
}
