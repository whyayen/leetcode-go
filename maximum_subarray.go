package main

func maxSubArray(nums []int) int {
	maxN := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum = sum + nums[i]
		} else {
			sum = nums[i]
		}

		if sum > maxN {
			maxN = sum
		}
	}

	return maxN
}
