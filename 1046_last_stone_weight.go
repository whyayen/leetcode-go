package main

import "container/heap"

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *IntHeap) Pop() any {
	n := len(*h)
	last := (*h)[n-1]
	*h = (*h)[:n-1]

	return last
}

func (h *IntHeap) Top() any {
	return (*h)[0]
}

func lastStoneWeight(stones []int) int {
	maxHeap := &IntHeap{}
	heap.Init(maxHeap)

	for _, stone := range stones {
		heap.Push(maxHeap, stone)
	}

	for maxHeap.Len() > 1 {
		y, x := heap.Pop(maxHeap).(int), heap.Pop(maxHeap).(int)

		if x != y {
			heap.Push(maxHeap, y-x)
		}
	}

	if maxHeap.Len() == 0 {
		return 0
	}
	return maxHeap.Top().(int)
}
