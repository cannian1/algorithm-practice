// 54. 螺旋矩阵
// https://leetcode.cn/problems/spiral-matrix/

package matrix

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, cols := len(matrix), len(matrix[0])
	result := make([]int, 0, rows*cols)

	// 定义四个边界
	top, bottom, left, right := 0, rows-1, 0, cols-1

	for top <= bottom && left <= right {
		// 从左到右遍历顶部行
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++

		// 从上到下遍历右侧列
		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right--

		// 确保还有行和列需要遍历
		if top <= bottom {
			// 从右到左遍历底部行
			for col := right; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
			bottom--
		}

		if left <= right {
			// 从下到上遍历左侧列
			for row := bottom; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++
		}
	}

	return result
}
