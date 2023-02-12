package main

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// TODO: review needed

type LRUCache struct {
	size     int
	capacity int
	cache    map[int]*DListNode
	head     *DListNode
	tail     *DListNode
}

type DListNode struct {
	key   int
	value int
	prev  *DListNode
	next  *DListNode
}

func NewDListNode(key, value int) *DListNode {
	return &DListNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DListNode),
		head:     NewDListNode(0, 0),
		tail:     NewDListNode(0, 0),
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.moveToHead(v)
		return v.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		v.value = value
		this.moveToHead(v)
	} else {
		node := NewDListNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	}
}

func (this *LRUCache) addToHead(node *DListNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DListNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DListNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
