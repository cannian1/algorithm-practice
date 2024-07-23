// 53.最大子数组和
// https://leetcode.cn/problems/maximum-subarray/

package array

func maxSubArray(nums []int) int {
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		// 在原数组上暂存局部最大子数组的结果，前一个元素与当前元素相加只要比当前元素大就覆盖当前元素（负数就直接被覆盖掉了）
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		// 与目前为止遇到的全局最优解比较
		maxSum = max(maxSum, nums[i])
	}
	return maxSum
}
