package astar

import (
	"container/heap"
)

type astar struct {
	grid [][]int
	rows int
	cols int
	h    Heuristic
}

func New(grid [][]int, h Heuristic) *astar {
	return &astar{
		grid: grid,
		h:    h,
		rows: len(grid),
		cols: len(grid[0]),
	}
}

func (a *astar) isWalkable(row, col int) bool {
	return row >= 0 && row < a.rows && col >= 0 && col < a.cols && a.grid[row][col] == 0
}

func (a *astar) getNeighbors(p Coords) []Coords {
	neighbors := make([]Coords, 0)
	row := p.Y
	col := p.X
	orthogonal := [][]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}
	for _, d := range orthogonal {
		newRow := row + d[1]
		newCol := col + d[0]
		if a.isWalkable(newRow, newCol) {
			neighbors = append(neighbors, Coords{newCol, newRow})
		}
	}

	diagonal := [][]int{
		{-1, -1},
		{-1, 1},
		{1, 1},
		{1, -1},
	}

	for _, d := range diagonal {
		newRow := row + d[1]
		newCol := col + d[0]

		if a.isWalkable(newRow, newCol) {
			//i want to move diagonally only if both the orthogonals are walkable
			if a.isWalkable(newRow, col) && a.isWalkable(row, newCol) {
				neighbors = append(neighbors, Coords{newCol, newRow})
			}
		}
	}

	return neighbors
}

func (a *astar) FindPath(start, end Coords) []Coords {

	openSet := make(PriorityQueue, 0)
	openMap := make(map[Coords]*node)
	closedSet := make(map[Coords]struct{})

	startNode := &node{
		position: start,
		g_cost:   0,
		h_cost:   a.h.estimate(start, end),
	}

	heap.Init(&openSet)
	heap.Push(&openSet, startNode)

	openMap[start] = startNode

	for openSet.Len() > 0 {
		current := heap.Pop(&openSet).(*node) //openSet.Pop().(*node)

		if current.position == end {
			return reconstructPath(current)
		}

		closedSet[current.position] = struct{}{}

		for _, neighborPos := range a.getNeighbors(current.position) {
			_, found := closedSet[neighborPos]
			if found {
				continue
			}
			tentativeG := current.g_cost + 1.0

			neighbor, found := openMap[neighborPos]

			if !found {
				neighbor = &node{
					position: neighborPos,
					g_cost:   tentativeG,
					h_cost:   a.h.estimate(neighborPos, end),
					parent:   current,
				}
				neighbor.f_cost = neighbor.g_cost + neighbor.h_cost
				openMap[neighborPos] = neighbor
				heap.Push(&openSet, neighbor)
				continue
			}

			if tentativeG < neighbor.g_cost {
				neighbor.g_cost = tentativeG
				neighbor.f_cost = neighbor.g_cost + neighbor.h_cost
				heap.Fix(&openSet, neighbor.index)
			}
		}
	}
	return nil
}

func reconstructPath(node *node) []Coords {
	path := make([]Coords, 0)
	for node != nil {
		path = append(path, node.position)
		node = node.parent
	}
	return reverse(path)
}

func reverse[T any](list []T) []T {
	for i, j := 0, len(list)-1; i < j; {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
	return list
}
