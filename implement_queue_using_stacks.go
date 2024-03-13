package main

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (this *Stack) push(num int) {
	this.data = append(this.data, num)
}

func (this *Stack) pop() int {
	if len(this.data) == 0 {
		panic("Invalid Operation")
	}

	pop := this.data[len(this.data)-1]
	this.data = this.data[0 : len(this.data)-1]
	return pop
}

func (this *Stack) isEmpty() bool {
	return len(this.data) == 0
}

func (this *Stack) peek() int {
	if len(this.data) == 0 {
		panic("Invalid Operation")
	}

	return this.data[len(this.data)-1]
}

type MyQueue struct {
	inStack  *Stack
	outStack *Stack
}

func Constructor() MyQueue {
	return MyQueue{inStack: NewStack(), outStack: NewStack()}
}

func (this *MyQueue) Push(x int) {
	this.inStack.push(x)
}

func (this *MyQueue) Pop() int {
	for !this.inStack.isEmpty() {
		pop := this.inStack.pop()
		this.outStack.push(pop)
	}

	shouldPop := this.outStack.pop()
	for !this.outStack.isEmpty() {
		pop := this.outStack.pop()
		this.inStack.push(pop)
	}
	return shouldPop
}

func (this *MyQueue) Peek() int {
	return this.inStack.data[0]
}

func (this *MyQueue) Empty() bool {
	return this.inStack.isEmpty() && this.outStack.isEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
