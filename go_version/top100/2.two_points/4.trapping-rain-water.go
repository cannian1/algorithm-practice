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

// 单调栈解法
func trap2(height []int) int {
	var stack []int // 栈，用来存储柱子的下标
	water := 0      // 用来累计接到的水量
	n := len(height)

	for i := range n {
		// 如果当前柱子比栈顶柱子高，那么可以计算接到的水量
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]   // 出栈
			stack = stack[:len(stack)-1] // 弹出栈顶元素

			if len(stack) > 0 {
				// 计算宽度：当前柱子与栈顶元素之间的距离
				width := i - stack[len(stack)-1] - 1
				// 计算水量：最小的高度减去出栈柱子的高度
				heightDiff := min(height[stack[len(stack)-1]], height[i]) - height[top]
				water += width * heightDiff
			}
		}
		// 当前柱子入栈
		stack = append(stack, i)
	}

	return water
}
