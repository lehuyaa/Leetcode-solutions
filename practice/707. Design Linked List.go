package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	start *Node
	end   *Node
	size  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		start: nil,
		end:   nil,
		size:  0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.size {
		return -1
	}
	cur := this.start
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{val: val, next: nil}
	if this.start == nil {
		this.start = node
		this.end = node
	}
	node.next = this.start
	this.start = node
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{val: val, next: nil}
	if this.end == nil {
		this.start = node
		this.end = node
	}
	this.end.next = node
	this.end = node
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		this.size++
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		this.size++
		return
	}
	cur := this.start
	for i := 0; i < index-1; i++ {
		cur = cur.next
	}
	node := &Node{val: val, next: nil}
	node.next = cur.next
	cur.next = node
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.size {
		return
	}
	if index == 0 {
		this.start = this.start.next
		this.size--
		if this.start == nil {
			this.end = nil
		}
		return
	}
	cur := this.start
	for i := 0; i < index-1; i++ {
		cur = cur.next
	}
	cur.next = cur.next.next
	if cur.next == nil {
		this.end = cur
	}
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	var obj MyLinkedList
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	fmt.Println(obj.Get(1))
	obj.DeleteAtIndex(1)
	fmt.Println(obj.Get(1))
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
