package main

type MyCircularQueue struct {
	Head  int
	Size  int
	Count int
	Queue []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Head:  0,
		Size:  k,
		Count: 0,
		Queue: make([]int, k),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	// push value into the queue,
	// return true if enqueue successfully.
	// return false if queue is full.
	// queue index is 0, 1, 2,..., k-1, 0, 1, 2.
	// tailIndex = (head + count - 1) mod cap(k).

	// is full ?
	if this.Count == this.Size {
		return false
	}
	// insert at the tail index
	this.Count++
	id := (this.Head + this.Count - 1) % this.Size
	// fmt.Printf("push: %d\n", id)
	this.Queue[id] = value

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	// remove one element from head, head index shift 1.
	// return true if dequeue successfully.
	// return false if queue is empty.
	if this.Count == 0 {
		return false
	}

	this.Head = (this.Head + 1) % this.Size
	// fmt.Printf("pop: %d\n", this.Head)
	this.Count -= 1
	return true
}

func (this *MyCircularQueue) Front() int {
	// return value of head element in the queue.
	// return -1 if queue is empty.
	if this.Count == 0 {
		return -1
	}
	return this.Queue[this.Head]
}

func (this *MyCircularQueue) Rear() int {
	// return value of tail element in the queue.
	// return -1 if queue is empty.
	if this.Count == 0 {
		return -1
	}
	id := (this.Head + this.Count - 1) % this.Size
	return this.Queue[id]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Count == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Count == this.Size
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
