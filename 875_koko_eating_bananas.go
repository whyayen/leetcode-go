package main

import "math"

/* 解法思路：
 * 找最大的 pile，然後做 Binary Search 去找最小 canEat 的 k
 * canEat 就是把每個 pile / k 並且無條件進位，代表該 pile 最少要耗費的 hours
 * 把每個 pile 耗費時間加總有 <= h 就代表 canEat
 * 然後從 1 ~ max pile 做 Binary Search，例如：
 * piles = [3,6,7,11], h = 8
 * k:      1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
 * canEat: F, F, F, T, T, T, T, T, T, T, T
 * mid 找尋過程會是 6(T) => 3(F) => 5(T) => 4(T) 此時 low == high 結束
 */
func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, pile := range piles {
		maxPile = max(maxPile, pile)
	}

	low, high := 1, maxPile
	for low < high {
		mid := low + (high-low)/2

		if canEat(piles, h, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func canEat(piles []int, h int, k int) bool {
	total := 0

	for _, pile := range piles {
		total = total + int(math.Ceil(float64(pile)/float64(k)))
	}

	return total <= h
}
