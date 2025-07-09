package leetcode_go

const DefaultCost = -1

type Item struct {
	source int
	target int
	weight int
}

type PriorityQueue []*Item

func (pq *PriorityQueue) Len() int { return len(*pq) }
func (pq *PriorityQueue) Less(i, j int) bool { return (*pq)[i].weight < (*pq)[j].weight }
func (pq *PriorityQueue) Swap(i, j int) { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }
func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Item))
}
func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	old := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return old
}

func costOfItem(v int) int {
	if v == DefaultCost {
		return math.MaxInt
	}

	return v
}

func networkDelayTime(times [][]int, n int, k int) int {
	pq := &PriorityQueue{}
	heap.Init(pq)

	edges := make(map[int][]*Item)
	for _, time := range times {
		edges[time[0]] = append(edges[time[0]], &Item{
			source: time[0],
			target: time[1],
			weight: time[2],
		})
	}

	cost := make(map[int]int)
	for i := 1; i <= n; i++ {
		if i == k {
			cost[i] = 0
		} else {
			cost[i] = DefaultCost
		}
	}

	for _, e := range edges[k] {
		heap.Push(pq, &Item{
			source: e.source,
			target: e.target,
			weight: e.weight,
		})
	}

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		costOfSource := costOfItem(cost[item.source])
		costOfTarget := costOfItem(cost[item.target])

		if costOfSource + item.weight < costOfTarget {
			cost[item.target] = costOfSource + item.weight

			for _, e := range edges[item.target] {
				heap.Push(pq, &Item{
					source: e.source,
					target: e.target,
					weight: e.weight,
				})
			}
		}
	}

	maxCost := 0
	for _, v := range cost {
		if v == DefaultCost {
			return -1
		}

		maxCost = max(maxCost, v)
	}

	return maxCost
}
