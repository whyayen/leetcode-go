package main

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int

	sort.Ints(nums)

	for i := 0; i <= len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			t := nums[i] + nums[left] + nums[right]

			if t == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if t > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return res
}
