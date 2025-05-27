package main

// f(i) = min(f[i-1], f[i-2]) + cost[i]

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	steps := make([]int, n)
	steps[0] = cost[0]
	steps[1] = cost[1]

	for idx := 2; idx < len(cost); idx++ {
		steps[idx] = min(steps[idx-1], steps[idx-2]) + cost[idx]
	}

	return min(steps[n-1], steps[n-2])
}
