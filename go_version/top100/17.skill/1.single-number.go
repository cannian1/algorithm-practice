// 136.只出现一次的数字
// https://leetcode.cn/problems/single-number/

package skill

// 两个相同的元素异或等于 0
func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}
