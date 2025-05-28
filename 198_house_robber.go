package main

// f(n) = f(n - 2) + nums[n]
// 要特別注意的測資：2,1,1,2

func rob(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	houses := len(nums)
	result := make([]int, houses)
	for i, num := range nums {
		if i < 2 {
			result[i] = nums[i]
		} else if i == 2 {
			result[i] = result[i-2] + num
		} else {
			result[i] = max(result[i-2]+num, result[i-3]+num)
		}
	}

	return max(result[houses-2], result[houses-1])
}
