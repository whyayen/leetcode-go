package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	lowestPrice := prices[0]
	maxP := 0

	for _, p := range prices {
		if p < lowestPrice {
			lowestPrice = p
		}

		if (p - lowestPrice) > maxP {
			maxP = p - lowestPrice
		}
	}
	return maxP
}

func main() {
	p := []int{5, 2, 6, 1, 3}
	result := maxProfit(p)
	fmt.Printf("%+v \n", result)
}
