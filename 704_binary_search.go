package main

import "fmt"

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		// 之所以 mid 不是 (low + high) / 2 是因為，這樣 low 或 high 已經到 int 極限的時候，在底層的程式相加會有溢位
		// 而用 low + (high - low)/2 則會確保在這兩者範圍中間，且其算式是等價的
		// 因爲可以推導為 (2 * low) / 2 + (high / 2) - (low / 2) => (2 * low + high - low) / 2 => (low + high) / 2
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	n := []int{-1, 0, 3, 5, 9, 12}
	r := search(n, 9)
	fmt.Printf("found: %+v \n", r)
}
