// 15. 三数之和
// https://leetcode.cn/problems/3sum/

package two_points

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	result := make([][]int, 0)
	for i := range nums {
		if nums[i] > 0 {
			return result
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return result
}
