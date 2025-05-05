package main

import "slices"

// 思路：像 78. Subsets 使用 backtracking 演算法，分別把數字 include & exclude
// 差異：
//    1. 因為一個 num 可以重複累加，且多了一個加總要 == target 的條件
// 所以：
//    1. 先將 candidates 排序，由大到小
//    2. 先對大的 number 一直重複累加
//    3. 當 idx 超過 candidates 大小，就表示把所有狀況都走完了，所以 return
//    4. 當 sum 已經超過 target 表示當前的數字已經加爆了，所以 return 讓其 backtrack 往下一個數字邁進
//    5. 當 sum == target 表示找到符合的組合，所以紀錄結果後 return 讓其 backtrack 找下一組和

func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	slices.Reverse(candidates)

	var explore func(int)
	var curr []int
	var sum int
	var result [][]int

	explore = func(idx int) {
		if sum == target {
			subset := make([]int, len(curr))
			copy(subset, curr)
			result = append(result, subset)
			return
		} else if sum > target || idx >= len(candidates) {
			return
		}

		curr = append(curr, candidates[idx])
		sum += candidates[idx]
		explore(idx)

		sum -= candidates[idx]
		curr = curr[:len(curr)-1]
		explore(idx + 1)
	}

	explore(0)
	return result
}
