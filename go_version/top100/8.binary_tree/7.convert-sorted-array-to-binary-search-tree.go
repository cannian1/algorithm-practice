// 108. 将有序数组转换为二叉搜索树
// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/

package binary_tree

// O(n) O(n)
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}
