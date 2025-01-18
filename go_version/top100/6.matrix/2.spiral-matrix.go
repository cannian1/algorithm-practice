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
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// 从上到下遍历右侧列
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// 确保还有行和列需要遍历
		if top <= bottom {
			// 从右到左遍历底部行
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			// 从下到上遍历左侧列
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}
