package main

func canConstruct(ransomNote string, magazine string) bool {
	chars := map[rune]int{}

	for _, ch := range ransomNote {
		_, isExist := chars[ch]

		if isExist {
			chars[ch] += 1
		} else {
			chars[ch] = 1
		}
	}

	for _, ma := range magazine {
		_, isExist := chars[ma]

		if isExist {
			chars[ma] -= 1
		}
	}

	for _, v := range chars {
		// 用 > 0 是因為可能 ransomeNote 是 ab
		// magazine 是 xyakbcababab
		// 因為 ransomeNote 用過了，所以 v 可能是負數的
		if v > 0 {
			return false
		}
	}

	return true
}

func main() {
	canConstruct("aa", "aab")
}
