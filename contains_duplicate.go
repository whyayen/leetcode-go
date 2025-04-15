package main

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] == nums[i+1] {
				return true
			}

			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				sorted = false
			}
		}
	}

	return false
}
