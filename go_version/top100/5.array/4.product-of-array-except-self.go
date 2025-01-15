// 238.除自身以外数组的乘积
// https://leetcode.cn/problems/product-of-array-except-self/

package array

// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// 初始化前缀积和后缀积
	// 遍历数组，将当前位置的值初始化为左侧所有元素的积
	prefix := 1 // 左侧的累积乘积
	for i := 0; i < n; i++ {
		answer[i] = prefix
		prefix *= nums[i]
	}

	// 计算后缀积并更新结果数组
	// 反向遍历数组，同时将当前结果乘以右侧所有元素的积
	suffix := 1 // 右侧的累积乘积
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix
		suffix *= nums[i]
	}

	return answer
}
