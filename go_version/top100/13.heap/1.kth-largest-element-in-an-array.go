// 215.数组中的第K个最大元素
// https://leetcode.cn/problems/kth-largest-element-in-an-array/

package heap

// 1.堆排序做法
func findKthLargest(nums []int, k int) int {
	hp := MakeHeap(nums)
	for i := 0; i < k-1; i++ {
		hp.Pop()
	}
	return hp.Pop()
}

type Heap struct {
	inner []int
}

func MakeHeap(nums []int) Heap {
	n := len(nums)
	for i := (n - 1) / 2; i >= 0; i-- {
		down(nums, i, n)
	}
	return Heap{nums}
}

func down(nums []int, root, end int) {
	child := root*2 + 1
	for child < end {
		if child+1 < end && nums[child+1] > nums[child] {
			child++
		}
		if nums[root] > nums[child] {
			break
		}
		nums[root], nums[child] = nums[child], nums[root]
		root = child
		child = root*2 + 1
	}
}

func (h *Heap) Pop() (res int) {
	ln := len(h.inner)
	if ln > 0 {
		res = h.inner[0]
	}
	h.inner[0] = h.inner[ln-1]
	h.inner = h.inner[:ln-1]
	down(h.inner, 0, ln-1)
	return
}

// 2. 计数排序做法
func findKthLargest2(nums []int, k int) int {
	var counts [20001]int
	for _, num := range nums {
		counts[num+10000]++
	}

	// 找到第 k 大的元素
	n := len(nums) - k // 从最小的元素开始，向上找到第 k 大的元素需要跳过的元素个数
	for num, cnt := range counts {
		n -= cnt
		if n < 0 {
			return num - 10000
		}
	}
	return -1
}

// 3.快速选择
