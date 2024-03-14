package main

func containsDuplicate(nums []int) bool {
	h := map[int]int{}

	for _, v := range nums {
		_, isExist := h[v]

		if isExist {
			return true
		} else {
			h[v] = 1
		}
	}

	return false
}
