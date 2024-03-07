package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	pairs := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}
	shouldPush := "{[("
	shouldPop := "}])"
	var stack []string

	for _, char := range s {
		if strings.Contains(shouldPush, string(char)) {
			stack = append(stack, string(char))
		} else if strings.Contains(shouldPop, string(char)) {
			if len(stack) < 1 {
				return false
			}

			popStr := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			v, _ := pairs[string(char)]
			if v != popStr {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	valid := isValid("[")
	fmt.Printf("%+v\n", valid)
}
