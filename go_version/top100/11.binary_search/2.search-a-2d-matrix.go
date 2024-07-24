// 74.搜索二维矩阵
// https://leetcode.cn/problems/search-a-2d-matrix/

package binary_search

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := left + ((right - left) >> 1)
		r := mid / n
		c := mid % n

		if matrix[r][c] > target {
			right = mid - 1
		} else if matrix[r][c] > target {
			left = mid + 1
		} else {
			return true
		}
	}
	return false
}
