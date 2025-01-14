// 239.滑动窗口最大值
// https://leetcode.cn/problems/sliding-window-maximum

package substring

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
// 返回 滑动窗口中的最大值

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	// 双端队列，存储的是元素的下标
	deque := make([]int, 0)
	result := make([]int, 0, len(nums)-k+1)

	for i, v := range nums {
		// 如果队列头部的元素已经不在滑动窗口范围内，移除它
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		// 移除队列中所有比当前元素小的元素
		for len(deque) > 0 && nums[deque[len(deque)-1]] < v {
			deque = deque[:len(deque)-1]
		}

		// 将当前元素下标加入队列
		deque = append(deque, i)

		// 从第 k 个元素开始记录结果
		if i >= k-1 {
			// 由于队首到队尾单调递减，所以窗口最大值就是队首
			result = append(result, nums[deque[0]])
		}
	}

	return result
}
