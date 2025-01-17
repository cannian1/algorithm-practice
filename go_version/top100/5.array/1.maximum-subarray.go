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

// 变形：返回最大子数组
func maxSubArray1(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	maxSum := nums[0]
	currentSum := nums[0]
	start := 0
	end := 0
	tempStart := 0

	for i := 1; i < len(nums); i++ {
		// 如果当前元素比加上前面的和更大，重置当前子数组的起点
		if nums[i] > currentSum+nums[i] {
			currentSum = nums[i]
			tempStart = i
		} else {
			currentSum += nums[i]
		}

		// 更新全局最大和及对应的子数组边界
		if currentSum > maxSum {
			maxSum = currentSum
			start = tempStart
			end = i
		}
	}

	// 返回最大子数组
	return nums[start : end+1]
}
