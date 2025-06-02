package main

// 重點
//   1. 可以「買」和「賣」不限次數
//   2. 但在「賣出」股票的隔天不能在買，必須等一天

// 思路
//  1. 使用狀態轉移的 DP
//  2. 設定 3 個狀態
//     a. hold[i]: 第 i 天持有股票的最大收益（也就是手上還持有股票）
//     b. sold[i]: 第 i 天賣出股票的最大收益（今天賣掉股票）
//     c. rest[i]: 第 i 天是 cooldown 或什麼都沒做的最大收益
func maxProfit(prices []int) int {
	n := len(prices)
	hold := -prices[0] // 第一天買入的收益
	sold := 0          // 第一天還不能賣，所以是 0
	rest := 0          // 第一天還沒有前一次賣的收益，所以是 0

	for i := 1; i < n; i++ {
		prevHold := hold
		prevSold := sold
		prevRest := rest

		hold = max(prevHold, prevRest-prices[i]) // 為前一天就持有，或之前的收益 - 當天購入的成本
		sold = prevHold + prices[i]              // 今天賣，一定是昨天有持股
		rest = max(prevRest, prevSold)           // 今天 cooldown，要嘛昨天就 cooldown，要嘛昨天賣了
	}

	return max(rest, sold)
}
