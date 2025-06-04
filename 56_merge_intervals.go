package main

import "sort"

// 參考：https://blog.techbridge.cc/2020/01/16/leetcode-%E5%88%B7%E9%A1%8C-pattern-merge-intervals/

func merge(intervals [][]int) [][]int {
	// 因為 intervals 沒排序過，所以我們照 intervals 的起始時間（intervals[i][0]）排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	prevInterval := intervals[0]

	i := 1
	for i < len(intervals) {
		// 如果前一個 interval 的結束時間 < 當前 interval 的起始時間，表示沒重疊
		// 直接 insert 到答案
		if prevInterval[1] < intervals[i][0] {
			res = append(res, prevInterval)
			prevInterval = intervals[i]
		} else {
			// 否則表示有重疊，要算重疊時間
			prevInterval[0] = min(prevInterval[0], intervals[i][0])
			prevInterval[1] = max(prevInterval[1], intervals[i][1])
		}

		i++
	}
	res = append(res, prevInterval)

	return res
}
