package main

import (
	"slices"
	"strconv"
)

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (this *Stack) Push(num int) {
	this.data = append(this.data, num)
}

func (this *Stack) Pop() int {
	if len(this.data) == 0 {
		panic("Invalid Operation")
	}

	pop := this.data[len(this.data)-1]
	this.data = this.data[0 : len(this.data)-1]
	return pop
}

func evalRPN(tokens []string) int {
	stack := NewStack()
	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	for _, token := range tokens {
		if opFunc, ok := ops[token]; ok {
			b := stack.Pop()
			a := stack.Pop()
			stack.Push(opFunc(a, b))
		} else if n, err := strconv.Atoi(token); err == nil {
			stack.Push(n)
		}
	}

	return stack.pop()
}
