package main

func climbStairs(n int) int {
	h := map[int]int{}

	for i := 0; i <= n; i++ {
		if i <= 1 {
			h[i] = 1
			continue
		}

		h[i] = h[i-2] + h[i-1]
	}

	return h[n]
}

func main() {
	climbStairs(5)
}
