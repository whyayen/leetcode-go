package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMaps := make(map[int][]int)

	for i, num := range nums {
		diff := target - num
		v, isExist := numMaps[diff]

		if !isExist {
			numMaps[num] = append(numMaps[num], i)
		} else {
			return []int{v[0], i}
		}
	}

	return nil
}

func main() {
	pos := twoSum([]int{7, 2, 3, 2, 5, 4}, 5)
	fmt.Printf("%+v\n", pos)
}
