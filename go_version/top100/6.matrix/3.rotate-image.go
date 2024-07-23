// 48. 旋转图像
// https://leetcode.cn/problems/rotate-image/

package matrix

// 先水平翻转，然后沿主对角线翻转
func rotate(matrix [][]int) {
	n := len(matrix)

	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}

	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
