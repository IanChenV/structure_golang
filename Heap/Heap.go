package structure

type IHeap []int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	size := len(old)
	x := old[size-1]
	*h = old[:size-1]
	return x
}
