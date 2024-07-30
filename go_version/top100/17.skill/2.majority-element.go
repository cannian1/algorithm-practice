// 169.多数元素
// https://leetcode.cn/problems/majority-element/

package skill

// 摩尔投票法 O(n),O(1)
func majorityElement(nums []int) int {
	count := 0 // 投票计数
	major := 0 // 当前候选人

	for i := 0; i < len(nums); i++ {
		// 如果票数都被抵消，就把当前元素设为候选人
		if count == 0 {
			major = nums[i]
		}

		// 如果遍历到的对象和当前候选人一致，票数增加
		// 否则票数减少
		if major == nums[i] {
			count++
		} else {
			count--
		}
	}
	return major
}
