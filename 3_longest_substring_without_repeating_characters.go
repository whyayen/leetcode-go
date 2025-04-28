package main

func lengthOfLongestSubstring(s string) int {
	ans := 0
	i := 0
	mp := map[string]int{}

	for j, c := range s {
		v, isExist := mp[string(c)]
		if isExist {
			i = max(i, v)
		}

		ans = max(ans, j-i+1)
		mp[string(c)] = j + 1
	}

	return ans
}
