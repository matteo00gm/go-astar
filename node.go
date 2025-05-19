package astar

//A node in the graph with position and cost information.
type node struct {
	position Coords
	g_cost   float64
	h_cost   float64
	f_cost   float64
	parent   *node
	index    int // The index of the item in the heap.
}
