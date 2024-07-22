// 1. 两数之和
// https://leetcode.cn/problems/two-sum/

package hash

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))
	for i, v := range nums {
		// 每遍历一个元素，看看哈希表里是否存在满足要求的目标数字
		if targetNumIdx, ok := mp[target-v]; ok {
			return []int{targetNumIdx, i}
		}
		mp[v] = i // 哈希表存储遍历过的元素和对应的索引
	}
	return nil
}
