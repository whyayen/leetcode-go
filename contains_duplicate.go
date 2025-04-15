package main

func containsDuplicate(nums []int) bool {
	nmap := make(map[int]int)

	for _, v := range nums {
		nmap[v]++

		if nmap[v] > 1 {
			return true
		}
	}

	return false
}
