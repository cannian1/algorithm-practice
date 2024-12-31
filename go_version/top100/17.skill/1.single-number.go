// 136.只出现一次的数字
// https://leetcode.cn/problems/single-number/

package skill

// 本题只有一个数字出现一次，其他数字都出现了两次

// 两个相同的元素异或等于 0
// 任何数和 0 做异或运算，结果仍然是原来的数
func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}
