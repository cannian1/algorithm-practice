// 146.LRU缓存
// https://leetcode.cn/problems/lru-cache

package linklist

import "container/list"

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

// 查询：当查询数据时，如果数据在缓存中，则将其移动到链表的头部。
// 插入：当插入新数据时，如果缓存未满，则直接将新数据插入链表头部，并在哈希表中记录；如果缓存已满，则淘汰链表尾部的数据，再插入新数据。
// 淘汰：当需要淘汰数据时，直接删除链表尾部的节点，并在哈希表中移除对应记录

func (this *LRUCache) Get(key int) int {
	if element, exists := this.cache[key]; exists {
		this.list.MoveToFront(element) // 标记为最近被使用
		return element.Value.(pair).value
	}
	return -1
}

/*
 1. 如果缓存的 key 存在，则将数据更新并移到链表头部
 2. 如果数据不存在
    2.1 如果长度大于等于 LRU 缓存的容量，移除最近最少使用的项（链表最后一个元素），并从缓存中删除

链表头部插入新的键值对
缓存存入新的链表节点
*/

func (this *LRUCache) Put(key int, value int) {
	if element, exists := this.cache[key]; exists {
		// 更新并标记为最近被使用
		element.Value = pair{key: key, value: value}
		this.list.MoveToFront(element)
	} else {
		if this.list.Len() >= this.capacity {
			// 移除最近最少使用的项
			lru := this.list.Back()
			if lru != nil {
				this.list.Remove(lru)
				delete(this.cache, lru.Value.(pair).key)
			}
		}
		// 插入新的键值对
		newItem := this.list.PushFront(pair{key: key, value: value})
		this.cache[key] = newItem
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
