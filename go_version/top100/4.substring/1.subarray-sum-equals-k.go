// 560.和为 k 的子数组
// https://leetcode.cn/problems/subarray-sum-equals-k

package substring

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数

func subarraySum(nums []int, k int) int {
	// 存储前缀和出现的次数
	prefixSumCount := make(map[int]int)
	prefixSumCount[0] = 1 // 初始状态，前缀和为 0 时的计数为 1

	// prefixSum 从数组起始到第 i 个元素的累计和
	count, prefixSum := 0, 0

	for _, num := range nums {
		// 更新当前前缀和
		prefixSum += num
		// 检查是否存在 prefixSum - k 的前缀
		if val, exists := prefixSumCount[prefixSum-k]; exists {
			count += val
		}
		// 更新当前前缀和的计数
		prefixSumCount[prefixSum]++
	}
	return count
}
