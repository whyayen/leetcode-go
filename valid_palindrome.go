package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	newStr := strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, ""))
	start := 0
	end := len(newStr) - 1

	for start <= end {
		if newStr[start] != newStr[end] {
			return false
		}

		start++
		end--
	}
	return true
}

func main() {
	valid := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Printf("%+v\n", valid)
}
