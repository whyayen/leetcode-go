package leetcode_go

type Edge struct {
	from int
	to int
	weight int
}

type PriorityQueue []*Edge

func (pq *PriorityQueue) Len() int { return len(*pq) }
func (pq *PriorityQueue) Less(i, j int) bool { return (*pq)[i].weight < (*pq)[j].weight }
func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}
func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Edge))
}
func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	old := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return old
}

// 用 Prim 演算法
// 因為每個點會連接其他所有點，所以會頂點少、邊多，適合 Prim 演算法，時間複雜 O(E + V log V)
// 如果用 Kruskal 則會是 O(E log E)
func minCostConnectPoints(points [][]int) int {
	pq := &PriorityQueue{}
	heap.Init(pq)

	visited := make(map[int]bool)
	totalCost := 0
	var calculateWeight func(i, j int) int
	calculateWeight = func(i, j int) int {
		return int(math.Abs(float64(points[i][0] - points[j][0])) + math.Abs(float64(points[i][1] - points[j][1])))
	}

	visited[0] = true
	for i := 1; i < len(points); i++ {
		heap.Push(pq, &Edge{
			from: 0,
			to: i,
			weight: calculateWeight(0, i),
		})
	}

	for pq.Len() > 0 {
		edge := heap.Pop(pq).(*Edge)
		if visited[edge.to] {
			continue
		}
		visited[edge.to] = true
		totalCost += edge.weight

		for i := 0; i < len(points); i++ {
			if i == edge.to || visited[i] {
				continue
			}

			heap.Push(pq, &Edge{
				from: edge.to,
				to: i,
				weight: calculateWeight(edge.to, i),
			})
		}
	}

	return totalCost
}

