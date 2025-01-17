// 73. 矩阵置零
// https://leetcode.cn/problems/set-matrix-zeroes/

package matrix

func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))

	for i, r := range matrix {
		for j := range r {
			// matrix[i][j] 和 r[j] 是一个东西
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] {
				r[j] = 0
			}
		}
	}
}

func setZeroes1(matrix [][]int) {
	row, col := false, false // 分别用来标记第一行和第一列是否需要设置为全零

	// 检查第一列是否包含零
	for i := range matrix {
		if matrix[i][0] == 0 {
			col = true
			break
		}
	}
	// 检查第一行是否包含零
	for j := range matrix[0] {
		if matrix[0][j] == 0 {
			row = true
			break
		}
	}

	// 从第二行和第二列开始遍历矩阵，使用第一行和第一列来标记需要设置为全零的行和列
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0 // 如果 matrix[i][j] 是零，则将第 i 行的第一个元素设为零
				matrix[0][j] = 0 // 如果 matrix[i][j] 是零，则将第 j 列的第一个元素设为零
			}
		}
	}
	// 根据第一行和第一列的标记，将需要设置为全零的行和列设为零
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0 // 如果第 i 行或第 j 列被标记为零，则将 matrix[i][j] 设为零
			}
		}
	}

	// 如果第一行需要被设置为全零，则执行以下操作
	if row {
		for j := range matrix[0] {
			matrix[0][j] = 0 // 将第一行的所有元素设为零
		}
	}
	if col {
		for i := range matrix {
			matrix[i][0] = 0 // 将第一列的所有元素设为零
		}
	}
}
