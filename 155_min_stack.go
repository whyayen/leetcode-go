package main

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, min(this.minStack[len(this.minStack)-1], val))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[0 : len(this.stack)-1]
	this.minStack = this.minStack[0 : len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
