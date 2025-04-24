package main

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	if n < 2 {
		return []int{}
	}

	x, y := 0, n-1

	for x < y {
		if numbers[x]+numbers[y] == target {
			return []int{x + 1, y + 1}
		}

		if numbers[x]+numbers[y] < target {
			x++
		} else {
			y--
		}
	}

	return []int{}
}
