package main

https://blog.techbridge.cc/2020/01/16/leetcode-%E5%88%B7%E9%A1%8C-pattern-merge-intervals/
// 解題思路：
//    當 intervals[i] 的 end < newInterval start 表示不會重疊，直接插入答案即可
//    上述動作做完後，就表示剩餘的 intervals 會有重疊，我們就持續做 merge intervals
//    所以繼續遍歷 intervals[i] 的 start <= newInterval end，持續拿到最小的 newInterval & 最大的 newInterval
//    merge intervals 結束後，把剩餘的 intervals 繼續插入到答案

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	i := 0

	// Skipping all intervals which ends before newInterval starts
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	// Keep merging until intervals[i] starts after newInterval
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}
	res = append(res, newInterval)

	// Push all left intervals into the result
	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}

	return res
}
