package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	smap := make(map[rune]int)

	for _, r := range s {
		smap[r]++
	}

	for _, r := range t {
		smap[r]--
	}

	for _, v := range smap {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	result := isAnagram("ab", "ab")
	fmt.Printf("%+v \n", result)
}
