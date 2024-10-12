package astar

// Len returns the number of elements in the priority queue
func (pq PriorityQueue) Len() int {
	return len(pq) // Length of the slice
}

// Less reports whether the element with index i has a lower priority than the element with index j
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].F < pq[j].F // Compare the F values (priority) of the nodes
}

// Swap exchanges the elements with indexes i and j
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i] // Swap nodes in the slice
}

// Push adds an element x to the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node)) // Append the new node to the priority queue
}

// Pop removes and returns the element with the highest priority (lowest F value) from the queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq       // Reference to the current slice
	n := len(old)    // Length of the slice
	x := old[n-1]    // Get the last element (highest priority)
	*pq = old[0 : n-1] // Resize the slice to remove the last element
	return x         // Return the popped element
}
