// 200.岛屿数量
// https://leetcode.cn/problems/number-of-islands

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。

package graph

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	islandCount := 0

	// 深度优先搜索 (DFS) 函数
	var dfs func(row, col int)
	dfs = func(i, j int) {
		// 如果超出边界或者当前格子是水（'0'），则返回
		if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] == '0' {
			return
		}

		// 将当前格子标记为已访问（'0'）
		grid[i][j] = '0'

		// 递归访问相邻的格子
		dfs(i-1, j) // 上
		dfs(i+1, j) // 下
		dfs(i, j-1) // 左
		dfs(i, j+1) // 右
	}

	// 遍历网格
	for i := range rows {
		for j := range cols {
			// 如果遇到陆地（'1'），启动一次DFS，并计数
			if grid[i][j] == '1' {
				islandCount++
				dfs(i, j)
			}
		}
	}

	return islandCount
}
