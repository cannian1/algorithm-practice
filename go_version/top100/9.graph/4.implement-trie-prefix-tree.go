// 208.实现前缀树
// https://leetcode.cn/problems/implement-trie-prefix-tree

package graph

type Trie struct {
	child [26]*Trie
	isEnd bool
}

// Constructor 构造函数
func Constructor() Trie {
	return Trie{}
}

// Insert 插入单词
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.child[ch] == nil {
			node.child[ch] = new(Trie)
		}
		node = node.child[ch]
	}
	node.isEnd = true
}

// Search 查找完整单词
func (t *Trie) Search(word string) bool {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.child[ch] == nil {
			return false
		}
		node = node.child[ch]
	}
	return node.isEnd
}

// StartsWith 查找前缀
func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.child[ch] == nil {
			return false
		}
		node = node.child[ch]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
