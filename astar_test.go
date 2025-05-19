package astar

import (
	"slices"
	"testing"
)

/*
to run a specific test: go test -run TestFindPath/"path not found"
to run them all: go test
*/

/*
	- using table driven tests for different implementations
	- TestFindPath will just test if the algorithm can find a path that goes from start to end (you don't need to provide the full path)
	- TestBestPath tests the best path (you have to provide the exact positions "expected")
*/

func TestFindPath(t *testing.T) {

	cases := []struct {
		name            string
		grid            [][]int
		expectPathFound bool
		start           Coords
		goal            Coords
	}{
		{
			name: "path not found",
			grid: [][]int{
				{0, 0, 0, 1, 0},
				{1, 1, 1, 1, 0},
				{0, 0, 0, 0, 0},
				{0, 1, 1, 1, 0},
				{0, 0, 0, 0, 0},
			},
			expectPathFound: false,
			start:           Coords{0, 0},
			goal:            Coords{4, 4},
		},
		{
			name: "5x5",
			grid: [][]int{
				{0, 0, 0, 1, 0},
				{1, 1, 0, 1, 0},
				{0, 0, 0, 0, 0},
				{0, 1, 1, 1, 0},
				{0, 0, 0, 0, 0},
			},
			expectPathFound: true,
			start:           Coords{0, 0},
			goal:            Coords{4, 4},
		},
		{
			name: "8x5",
			grid: [][]int{
				{0, 0, 0, 1, 0, 1, 1, 1},
				{1, 1, 0, 1, 0, 1, 1, 1},
				{0, 0, 0, 0, 0, 1, 0, 0},
				{0, 1, 1, 1, 0, 1, 0, 1},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			expectPathFound: true,
			start:           Coords{0, 0},
			goal:            Coords{4, 4},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			astar := New(test.grid, &EuclideanHeuristic{})
			got := astar.FindPath(test.start, test.goal)

			if test.expectPathFound {
				if len(got) == 0 {
					t.Errorf("Expected a non-empty path for case '%s', but got an empty slice", test.name)
				}
			} else {
				if len(got) > 0 {
					t.Errorf("Expected no path found for case '%s', but got a path of length %d", test.name, len(got))
				}
			}
		})
	}
}

func TestBestPath(t *testing.T) {

	cases := []struct {
		name     string
		grid     [][]int
		expected []Coords
		start    Coords
		goal     Coords
	}{
		{
			name: "5x5",
			grid: [][]int{
				{0, 0, 0, 1, 0},
				{1, 1, 0, 1, 0},
				{0, 0, 0, 0, 0},
				{0, 1, 1, 1, 0},
				{0, 0, 0, 0, 0},
			},
			expected: []Coords{
				{0, 0},
				{1, 0},
				{2, 0},
				{2, 1},
				{2, 2},
				{3, 2},
				{4, 2},
				{4, 3},
				{4, 4},
			},
			start: Coords{0, 0},
			goal:  Coords{4, 4},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			astar := New(test.grid, &EuclideanHeuristic{})
			got := astar.FindPath(test.start, test.goal)

			if !slices.Equal(got, test.expected) {
				t.Errorf("got %v, want %v", got, test.expected)
			}
		})
	}
}
