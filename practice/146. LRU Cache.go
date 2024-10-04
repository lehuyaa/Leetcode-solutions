package main

type Node struct {
	val  int
	next *Node
	prev *Node
	key  int
}

type LRUCache struct {
	cache map[int]*Node
	head  *Node
	tail  *Node
	cap   int
	size  int
}

func Constructor(capacity int) LRUCache {
	fake_head := &Node{val: -1, prev: nil, next: nil}
	fake_last := &Node{val: -1, prev: nil, next: nil}
	fake_head.next = fake_last
	fake_last.prev = fake_head
	return LRUCache{
		cache: make(map[int]*Node, capacity),
		size:  0,
		cap:   capacity,
		head:  fake_head,
		tail:  fake_last,
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.cache[key]
	if !ok {
		return -1
	}

	if this.head.next == v {
		return v.val
	}

	this.SwapFist(v)

	return v.val
}

func (this *LRUCache) SwapFist(v *Node) {
	first := this.head.next
	next := v.next
	prev := v.prev

	this.head.next = v
	v.prev = this.head
	prev.next = next
	next.prev = prev
	first.prev = v
	v.next = first
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		this.head.next.val = value
		return
	}

	var cur *Node
	if this.size >= this.cap {
		this.size--
		cur = this.tail.prev
		cur.prev.next = this.tail
		this.tail.prev = cur.prev
		delete(this.cache, cur.key)
	} else {
		cur = &Node{}
	}
	this.size++
	next := this.head.next
	cur.val = value
	cur.key = key
	cur.next = next
	cur.prev = this.head
	this.head.next = cur
	next.prev = cur
	this.cache[key] = cur
}
