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

func evalRPN(tokens []string) int {
	numStack := NewStack()
	operators := []string{"+", "-", "*", "/"}

	for _, token := range tokens {
		isOperator := slices.Contains(operators, token)

		if !isOperator {
			num, _ := strconv.Atoi(token)
			numStack.push(num)
			continue
		}

		num2, num1 := numStack.pop(), numStack.pop()
		tmpResult := 0

		switch token {
		case "+":
			tmpResult = num1 + num2
		case "-":
			tmpResult = num1 - num2
		case "*":
			tmpResult = num1 * num2
		case "/":
			tmpResult = num1 / num2
		}
		numStack.push(tmpResult)
	}

	return numStack.pop()
}
