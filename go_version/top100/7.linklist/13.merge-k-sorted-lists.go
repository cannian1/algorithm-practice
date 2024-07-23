package linklist

import "container/heap"

type hp []*ListNode

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(*ListNode))
}

func (h *hp) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := hp{}
	for _, node := range lists {
		if node != nil {
			h = append(h, node)
		}
	}
	heap.Init(&h)

	dummy := &ListNode{} // 哨兵节点，作为合并后链表头节点的前一个节点
	cur := dummy
	for len(h) > 0 { // 循环直到堆为空
		node := heap.Pop(&h).(*ListNode) // 剩余节点中的最小节点
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
		cur.Next = node // 合并到新链表中
		cur = cur.Next  // 准备合并下一个节点
	}
	return dummy.Next // 哨兵节点的下一个节点就是新链表的头节点
}
