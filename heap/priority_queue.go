package heap

import (
	"container/heap"
)

// Go native implementation: https://github.com/golang/go/blob/master/src/container/heap/heap.go

// Demonstrate heap usage with PriorityQueue implementation

// Item is an item in the priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	index    int    // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func NewPriorityQueue() heap.Interface {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return &pq
}

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].priority < (*pq)[j].priority
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

// Push adds x as last element of the PriorityQueue
func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	item.index = pq.Len()
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := pq.Len()
	lastItem := old[n-1]
	old[n-1] = nil      // don't stop the GC from reclaiming the item eventually
	lastItem.index = -1 // for safety
	*pq = old[0 : n-1]
	return lastItem
}
