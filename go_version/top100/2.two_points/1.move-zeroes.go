// 283. 移动零
// https://leetcode.cn/problems/move-zeroes/

package two_points

func moveZeroes(nums []int) {
	// 双指针，交换0与遍历过程中的非0元素
	k := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}
}
