package main

func majorityElement(nums []int) int {
	maxN := nums[0]
	times := 0

	for _, v := range nums {
		if maxN != v {
			times--
		} else {
			times++
		}

		if times < 0 {
			times = 0
			maxN = v
		}
	}

	return maxN
}

func main() {
	
}
