package main

// 計算二進制有幾個 1
// 可以採用 Brian Kernighan's Algorithm
// 算法：每次跟數字 - 1 做 And 運算，然後把 counter + 1
// 原理：讓 n 每次最右邊的 1 被消除
//
//	110100 & 110010 => 110000 (count = 1)
//	110000 & 101000 => 100000 (count = 2)
//	100000 & 010000 => 0 (count = 3)
func hammingWeight(n int) int {
	count := 0

	for n != 0 {
		n = n & (n - 1)
		count++
	}

	return count
}
