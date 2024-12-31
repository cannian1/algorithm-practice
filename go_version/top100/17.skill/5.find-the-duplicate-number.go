// 287.寻找重复数
// https://leetcode.cn/problems/find-the-duplicate-number/

package skill

// 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
// 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
// 不修改数组 nums 且只用常量级 O(1) 的额外空间。

// findDuplicate 使用 Floyd 判圈算法寻找重复数 （算法思想：142. 环形链表 II）
// O(n)、O(1)
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]

	for {
		slow = nums[slow]       // 慢指针每次走一步
		fast = nums[nums[fast]] // 快指针每次走两步

		if slow == fast { // 找到相遇点
			break
		}
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
