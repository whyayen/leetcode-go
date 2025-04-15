package main

func topKFrequent(nums []int, k int) []int {
	tmap := make(map[int]int)

	for _, n := range nums {
		tmap[n]++
	}

	frequents := make([][]int, len(tmap))
	count := 0

	for k, v := range tmap {
		if count == 0 {
			frequents[0] = []int{v, k}
		} else {
			j := count - 1
			for j >= 0 && frequents[j][0] < v {
				frequents[j+1] = frequents[j]
				j--
			}

			frequents[j+1] = []int{v, k}
		}

		count++
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = frequents[i][1]
	}

	return result
}
