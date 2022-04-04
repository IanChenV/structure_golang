package structue

import (
	"container/heap"
)

type entry struct {
	key      string
	priority int
	// index 是entry在heap中的索引号
	// entry 加入Priority Queue后，Priority会变化时，很有用
	// 如果 entry.priority 一直不变的话，可以删除index
	index int
}

type PQ []*entry

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[j].index = i
	pq[i].index = j
}

func (pq *PQ) Push(x interface{}) {
	temp := x.(*entry)
	temp.index = len(*pq)
	*pq = append(*pq, temp)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	size := len(old)
	x := old[size-1]
	*pq = old[:size-1]
	x.index = -1
	return x
}

func (pq *PQ) update(entry *entry, value string, priority int) {
	entry.key = value
	entry.priority = priority
	heap.Fix(pq, entry.index)
}
