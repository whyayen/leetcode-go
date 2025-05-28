package main

import (
	"fmt"
	"math"
	"slices"
)

// 從總額減去可兌換的硬幣面額，下去推導該數字 + 該面額硬幣 x 1，找最小就可以找出組合關係
// 把 f(0) 視為 0，< 0 的視為無限大兌換不了硬幣
// 11 => [1,2,5]
//   f(11) => min(f(10) + 1, f(9) + 1, f(6) + 1)
//      | f(10) => min(f(9) + 1, f(8) + 1, f(5) + 1)
//      | f(9) => min(f(8) + 1, f(7) + 1, f(4) + 1)
//      | f(8) => min(f(7) + 1, f(6) + 1, f(3) + 1)
//      | f(7) => min(f(6) + 1, f(5) + 1, f(2) + 1)
//      | f(6) => min(f(5) + 1, f(4) + 1, f(1) + 1)
//      | f(5) => min(f(4) + 1, f(3) + 1, f(0) + 1)
//      | f(4) => min(f(3) + 1, f(2) + 1, f(-1) + 1)
//      | f(3) => min(f(2) + 1, f(1) + 1, f(-2) + 1)
//      | f(2) => min(f(1) + 1, f(0) + 1, f(-3) + 1)
//      | f(1) => min(f(0) + 1, f(-1) + 1, f(-4) + 1)
//      | f(0) => 0
//      | f(-1 ~ -4) => 無限大

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
