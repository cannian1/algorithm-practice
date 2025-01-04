// 31.下一个排列
// https://leetcode.cn/problems/next-permutation/

package skill

// 将数字重排列为下一个字典序的排列
// 从右向左找到第一个较小的数（递减点）
// 如果某个元素比其右侧的元素小，说明可以通过交换和调整右侧元素，构造出比当前排列大的下一个字典序排列

// 从递减点向右，找到比它大的最小值（保证字典序最小）
// 交换两个数后，将递减点之后的部分升序排列

func nextPermutation(nums []int) {
	// 从右向左找到第一个较小的数 (递减点)
	n := len(nums)
	if n <= 1 {
		return
	}

	i := n - 2 // 如果从 n - 1 开始，右边没有可以比较的数
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		// 从右向左找到第一个比 nums[i] 大的数
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		// 交换 nums[i] 和 nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 反转 i+1 到末尾的部分
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
