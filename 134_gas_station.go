package main

func canCompleteCircuit(gas []int, cost []int) int {
	// gas    3,4,5,1,2
	// cost   5,4,3,2,1
	// gain  -2,0,2,-1,1

	totalGas := 0
	currGas := 0
	startAt := 0

	for i := range gas {
		gain := gas[i] - cost[i]
		totalGas += gain
		currGas += gain

		// 只要目前的 Gas < 0 表示原本的起點就不適合，直接改從 i+1 開始
		// 並把 currGas 歸零
		if currGas < 0 {
			startAt = i + 1
			currGas = 0
		}
	}

	// 如果 totalGas < 0 表示根本就沒有夠的油量跑完一圈
	if totalGas < 0 {
		return -1
	}

	return startAt
}
