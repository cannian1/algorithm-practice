// 56. 合并区间
// https://leetcode.cn/problems/merge-intervals/

package array

import "sort"

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// 按照左端点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		// 当前区间与上一个区间存在重叠，合并当前区间与上一个区间
		if res[len(res)-1][1] >= intervals[i][0] {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			// 当前区间与上一个区间不重叠
			res = append(res, intervals[i]) // 直接添加当前区间
		}
	}
	return res
}
