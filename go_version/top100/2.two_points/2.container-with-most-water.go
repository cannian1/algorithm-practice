//  11. 盛最多水的容器
// https://leetcode.cn/problems/container-with-most-water

package two_points

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0
	for left < right {
		if height[left] < height[right] {
			result = max(result, height[left]*(right-left))
			left++
		} else {
			result = max(result, height[right]*(right-left))
			right--
		}
	}
	return result
}
