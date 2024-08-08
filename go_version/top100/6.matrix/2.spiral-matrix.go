// 54. 螺旋矩阵
// https://leetcode.cn/problems/spiral-matrix/

package matrix

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, cols := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, cols-1, 0, rows-1
	result := make([]int, 0, rows*cols)

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		for i := top + 1; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		if left < right && top < bottom {
			for i := right - 1; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			for i := bottom - 1; i > top; i-- {
				result = append(result, matrix[i][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}
