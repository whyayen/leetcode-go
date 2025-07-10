package leetcode_go

import "sort"

// 思考一個場景
//
//	JFK -> SFO
//	JFK -> ATL
//	ATL -> JFK
//
// 如果用 directed graph 一般走訪的概念，先走 JFK -> SFO 就會走不回來
// 而這裡的 DFS 不是單純貪心的 DFS，而是基於 Hierholzer’s Algorithm 的概念，適用於「找歐拉路徑（Eulerian Path）」：
// 用完每一條邊一次且僅一次（也就是：每張機票要用一次）。
func findItinerary(tickets [][]string) []string {
	graph := make(map[string][]string)

	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		graph[from] = append(graph[from], to)
	}

	// 為了方便 pop 出字典序最小的目的地 → 先把每個列表反向排序
	for key := range graph {
		sort.Sort(sort.Reverse(sort.StringSlice(graph[key])))
	}

	var itinenary []string
	var dfs func(string)
	dfs = func(airport string) {
		for len(graph[airport]) > 0 {
			// 取出字典序最小的下一站（因為反向排好，所以從末端 pop）
			last := len(graph[airport]) - 1
			next := graph[airport][last]
			graph[airport] = graph[airport][:last]

			dfs(next)
		}
		// 當前沒有後續了，加入結果（逆序加入）
		itinenary = append(itinenary, airport)
	}
	dfs("JFK")

	lastIndex := len(itinenary) - 1
	// 返轉結果
	for i := 0; i < (lastIndex+1)/2; i++ {
		itinenary[i], itinenary[lastIndex-i] = itinenary[lastIndex-i], itinenary[i]
	}

	return itinenary
}
