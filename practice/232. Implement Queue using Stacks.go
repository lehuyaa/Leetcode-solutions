package main

type MyQueue struct {
	stack      []int
	emptyStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack:      []int{},
		emptyStack: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, x)
		return
	}

	for len(this.stack) > 0 {
		size := len(this.stack)
		top := this.stack[size-1]
		this.stack = this.stack[:size-1]
		this.emptyStack = append(this.emptyStack, top)
	}

	this.stack = append(this.stack, x)

	for len(this.emptyStack) > 0 {
		size := len(this.emptyStack)
		top := this.emptyStack[size-1]
		this.stack = append(this.stack, top)
		this.emptyStack = this.emptyStack[:size-1]
	}

	this.emptyStack = []int{}
}

func (this *MyQueue) Pop() int {
	if len(this.stack) == 0 {
		return 0
	}

	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	return top
}

func (this *MyQueue) Peek() int {
	if len(this.stack) == 0 {
		return 0
	}

	top := this.stack[len(this.stack)-1]
	return top
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
