package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sCharCount := map[string]int{}

	for i := 0; i < len(s); i++ {
		sKey := string(s[i])
		tKey := string(t[i])
		_, sVIsExist := sCharCount[sKey]

		if !sVIsExist {
			sCharCount[sKey] = 1
		} else {
			sCharCount[sKey]++
		}

		_, tVIsExist := sCharCount[tKey]
		if !tVIsExist {
			sCharCount[tKey] = -1
		} else {
			sCharCount[tKey]--
		}
	}

	for _, v := range sCharCount {
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
