// 118.杨辉三角
// https://leetcode.cn/problems/pascals-triangle

package dp

// 杨辉三角的每一排左对齐
// [1]
// [1,1]
// [1,2,1]
// [1,3,3,1]
// [1,4,6,4,1]

// 设要计算的二维数组是 c，计算方法如下
// 		每一排的第一个数和最后一个数都是 1，即 c[i][0]=c[i][i]=1
//		其余数字，等于左上方的数，加上正上方的数，即 c[i][j]=c[i−1][j−1]+c[i−1][j]。例如 4=1+3, 6=3+3 等

func generate(numRows int) [][]int {
	c := make([][]int, numRows)
	for i := range c {
		c[i] = make([]int, i+1)
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			// 左上方的数 + 正上方的数
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
	return c
}
