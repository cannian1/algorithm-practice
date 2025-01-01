// 84.柱状图中最大的矩形
// https://leetcode.cn/problems/largest-rectangle-in-histogram

package stack

// 单调栈解法

func largestRectangleArea(heights []int) int {
	// 创建一个栈，用于存储柱子的索引
	var stack []int
	maxArea := 0

	// 在柱子后面添加一个高度为 0 的柱子，简化计算
	// 确保在循环结束时，栈中所有未处理的柱子都会被弹出并计算面积
	heights = append(heights, 0)

	for i := range heights {
		// 当栈不为空且当前柱子的高度小于栈顶柱子的高度时
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]] // 获取栈顶柱子的高度
			stack = stack[:len(stack)-1]      // 将栈顶柱子弹出

			// 计算宽度
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			// 更新最大面积
			maxArea = max(maxArea, h*width)
		}
		// 将当前柱子的索引压入栈中
		stack = append(stack, i)
	}

	return maxArea
}
