// 347. 前 K 个高频元素
// https://leetcode.cn/problems/top-k-frequent-elements/

package heap

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	mp := map[int]int{}

	for _, num := range nums {
		mp[num]++
	}

	hp := &IHeap{}
	heap.Init(hp)

	// 所有元素入堆，堆的长度为 k
	for key, val := range mp {
		heap.Push(hp, [2]int{key, val})
		if hp.Len() > k {
			heap.Pop(hp)
		}
	}
	res := make([]int, k)
	for i := range k {
		res[k-i-1] = heap.Pop(hp).([2]int)[0]
	}
	return res
}

type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
