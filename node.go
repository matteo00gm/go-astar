package astar

// A node in the graph with position and cost information.
type node struct {
	position Coords
	g_cost   float32
	h_cost   float32
	f_cost   float32
	parent   *node
	index    int // The index of the item in the heap.
}
