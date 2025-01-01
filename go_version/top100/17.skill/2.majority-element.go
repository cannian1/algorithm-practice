// 169.多数元素
// https://leetcode.cn/problems/majority-element/

package skill

import "errors"

// 这题假定数组必然存在 多数元素
// 如果不存在多数元素，参考下面第二种方法做验证

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

func majorityElement2(nums []int) (int, error) {
	major := 0
	count := 0

	for i := range nums {
		if count == 0 {
			major = nums[i]
		}

		if major == nums[i] {
			count++
		} else {
			count--
		}
	}

	// 第二遍遍历：验证候选元素是否为多数元素
	count = 0
	for _, num := range nums {
		if num == major {
			count++
		}
	}

	if count > len(nums)/2 {
		return major, nil
	}

	return 0, errors.New("不存在多数元素")
}
