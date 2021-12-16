package main

import "container/heap"

/** copied entirely from https://pkg.go.dev/container/heap@go1.17.5 */

// An Item is something we manage in a priority queue.
type Item struct {
	// The value of the item; arbitrary.
	value *Cell
	// The priority of the item in the queue.
	priority int
	// The index is needed by update and is maintained by the heap.Interface methods.
	// The index of the item in the heap.
	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the LOWEST priority so we use less than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(cell *Cell) {
	// find item
	for _, item := range *pq {
		if item.value == cell {
			item.priority = cell.distance
			heap.Fix(pq, item.index)

			break
		}
		// else: uh oh
	}
}
