package main

import (
	"fmt"
	"math"
	"slices"
)

func coinChange(coins []int, amount int) int {
	table := map[int]int{}
	result := dp(coins, amount, table)
	fmt.Printf("%+v\n", table)

	if result == math.MaxInt {
		return -1
	}
	return result
}

func dp(coins []int, amount int, table map[int]int) int {
	if amount < 0 {
		return math.MaxInt
	}

	if amount == 0 {
		return 0
	}

	v, amountIsExist := table[amount]
	if amountIsExist {
		return v
	}

	tmp := []int{}
	for _, coin := range coins {
		r := dp(coins, amount-coin, table)

		if r == math.MaxInt {
			tmp = append(tmp, r)
		} else {
			tmp = append(tmp, r+1)
		}
	}
	fmt.Printf("dp: %+v, %+v\n", amount, tmp)
	table[amount] = slices.Min(tmp)
	return table[amount]
}

func main() {
	coinChange([]int{1, 2, 5}, 11)
}
