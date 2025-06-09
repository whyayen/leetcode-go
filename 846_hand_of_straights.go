package main

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	cards := len(hand)
	if cards%groupSize != 0 {
		return false
	}

	sort.Ints(hand)
	cardCount := make(map[int]int)

	for _, card := range hand {
		cardCount[card]++
	}

	for _, card := range hand {
		// 不是牌組的頭，所以已經被前面扣掉過了，直接跳過
		if cardCount[card] == 0 {
			continue
		}

		for i := 0; i < groupSize; i++ {
			if cardCount[card+i] == 0 {
				return false
			}
			cardCount[card+i]--
		}
	}

	return true
}
