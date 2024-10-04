package main

type MyStack struct {
	queue      []int
	emptyQueue []int
}

func Constructor() MyStack {
	return MyStack{
		queue:      []int{},
		emptyQueue: []int{},
	}
}

func (this *MyStack) Push(x int) {
	if len(this.queue) == 0 {
		this.queue = append(this.queue, x)
		return
	}

	for len(this.queue) > 0 {
		top := this.queue[0]
		this.queue = this.queue[1:]
		this.emptyQueue = append(this.emptyQueue, top)
	}

	this.queue = append(this.queue, x)
	for len(this.emptyQueue) > 0 {
		topEmpty := this.emptyQueue[0]
		this.emptyQueue = this.emptyQueue[1:]
		this.queue = append(this.queue, topEmpty)
	}

	this.emptyQueue = []int{}
}

func (this *MyStack) Pop() int {
	if len(this.queue) == 0 {
		return 0
	}
	top := this.queue[0]
	this.queue = this.queue[1:]
	return top
}

func (this *MyStack) Top() int {
	if len(this.queue) == 0 {
		return 0
	}
	top := this.queue[0]
	return top
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
