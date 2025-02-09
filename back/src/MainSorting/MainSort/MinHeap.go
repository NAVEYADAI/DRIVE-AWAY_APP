package MainSort

import "container/heap"

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Grate < h[j].Grate }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) isEmpty() bool {
	return len(*h) == 0
}

func (h *MinHeap) Push(x interface{}) {
	item := x.(Item)
	*h = append(*h, item)
	heap.Fix(h, h.Len()-1)
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]

	return item
}
