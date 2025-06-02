package main

// 運用向右位移一位的結果來 + 最後一位是 1 or 0
func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 1; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}

	return ans
}
