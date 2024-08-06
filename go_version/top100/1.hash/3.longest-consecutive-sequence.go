// 128.最长连续序列
// https://leetcode.cn/problems/longest-consecutive-sequence

package hash

func longestConsecutive(nums []int) int {
	result := 0                             // 记录最长连续序列长度
	numSet := make(map[int]bool, len(nums)) // nums 中所有元素放入集合

	for _, num := range nums {
		numSet[num] = true
	}

	for _, num := range nums {
		// 这个数字没有前驱数才进入下面的逻辑
		if !numSet[num-1] {
			seqLen := 1 // 初始长度为 1
			// 统计连续序列长度
			for numSet[num+1] {
				seqLen += 1
				num += 1
			}
			result = max(result, seqLen)
		}
	}
	return result
}
