package astar

// Len returns the number of elements in the priority queue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less reports whether the element with index i has a lower priority than the element with index j
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].F < pq[j].F
}

// Swap exchanges the elements with indexes i and j
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push adds an element x to the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

// Pop removes and returns the element with the highest priority (lowest F value) from the queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
