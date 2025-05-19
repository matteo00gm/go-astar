package astar

type PriorityQueue []*node

func (pq PriorityQueue) Len() int { return len(pq) }

// Less determines the priority. For A*, lower f_cost is higher priority.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].f_cost < pq[j].f_cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	node := x.(*node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil  // avoid memory leak
	node.index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}
