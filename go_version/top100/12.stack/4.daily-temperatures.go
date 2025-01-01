// 739.每日温度
// https://leetcode.cn/problems/daily-temperatures/

package stack

// 单调栈，适合求解当前元素左边或者右边第一个比当前元素小或者大的元素
//
// 单调栈的作用：用于记录之前遍历过的元素，这样在遍历当前元素时才能知道当前元素是不是第一个比之前遍历过的元素大或者小的元素
//
// 单调栈如何选择递增还是递减？（从栈口往栈底看）
// 递增：当前元素后面第一个比它大的元素
// 递减：当前元素后面第一个比它小的元素

// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，
// 其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，该位置用 0 来代替。

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	// 单调栈，用以存放遍历数组的索引
	var stack []int
	for i, v := range temperatures {
		// 栈不空，且当前遍历元素 v 破坏了栈的单调性
		// 如果要入栈的元素比栈顶元素大，更新此时栈顶元素下标对应的结果，然后弹出去，如果还比栈顶元素大，还一直弹，直到没有比他大的了，然后入栈
		for len(stack) != 0 && v > temperatures[stack[len(stack)-1]] {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 当前遍历元素下标 - 当前单调栈栈口存放的元素下标
			res[top] = i - top
		}
		stack = append(stack, i)
	}
	return res
}
