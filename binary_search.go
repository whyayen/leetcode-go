package main

import "fmt"

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2

		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

func main() {
	n := []int{-1, 0, 3, 5, 9, 12}
	r := search(n, 9)
	fmt.Printf("found: %+v \n", r)
}
