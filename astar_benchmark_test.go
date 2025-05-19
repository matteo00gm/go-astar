package astar

import (
	"testing"
)

func BenchmarkFindPath(b *testing.B) {
	grid := [][]int{
		{0, 0, 0, 1, 0, 0, 0, 0},
		{1, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	start := Coords{X: 0, Y: 0}
	goal := Coords{X: 7, Y: 7}

	astar := New(grid, &EuclideanHeuristic{})

	// --- Benchmark Loop ---
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		astar.FindPath(start, goal)
	}
}
