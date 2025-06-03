package main

import "container/heap"

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
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

func (h *IntHeap) UpdateKthLargest(x any) {
	(*h)[0] = x.(int)
	heap.Fix(h, 0)
}

type KthLargest struct {
	minHeap *IntHeap
	k       int
}

func Constructor(k int, nums []int) KthLargest {
	minHeap := &IntHeap{}
	heap.Init(minHeap)
	obj := &KthLargest{
		k:       k,
		minHeap: minHeap,
	}

	for _, num := range nums {
		obj.push(num)
	}

	return *obj
}

func (this *KthLargest) Add(val int) int {
	this.push(val)
	return this.minHeap.Top().(int)
}

func (this *KthLargest) push(val int) {
	if this.minHeap.Len() < this.k {
		heap.Push(this.minHeap, val)
	} else if val > this.minHeap.Top().(int) {
		this.minHeap.UpdateKthLargest(val)
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
