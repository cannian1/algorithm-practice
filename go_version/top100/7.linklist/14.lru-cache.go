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

func (this *LRUCache) Get(key int) int {
	if element, exists := this.cache[key]; exists {
		this.list.MoveToFront(element) // 标记为最近被使用
		return element.Value.(pair).value
	}
	return -1
}

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
