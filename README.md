# Go A\* Pathfinding Implementation

This project provides a Go implementation of the A\* search algorithm for finding the shortest path on a 2D grid. It supports both orthogonal (up, down, left, right) and diagonal movements.

## Algorithm Overview

The A\* algorithm is an informed search algorithm that efficiently finds the shortest path between two points in a graph or grid. It uses a heuristic function to estimate the cost from the current node to the goal, guiding the search towards the target.

This implementation follows the standard A\* process, utilizing a priority queue to manage the set of open nodes (nodes to be evaluated) and a set to keep track of closed nodes (nodes already evaluated).

For a detailed understanding of the A\* algorithm, I highly recommend this article: [A\* Algorithm Explained](https://researchdatapod.com/a-star-algorithm/)

## Movement and Costs

Our implementation allows for movement to all 8 adjacent neighbors (4 orthogonal and 4 diagonal).

**Note that for a diagonal move to be valid, both of the adjacent orthogonal cells (those that share an edge with both the starting cell and the target diagonal cell) must also be walkable.**

**Important Note:** In this specific implementation, both orthogonal and diagonal movements are treated as having a uniform cost of **1.0**. This means the algorithm will find the path with the **minimum number of steps** rather than the minimum geometric distance, and it will tend to prefer diagonal movements when available as they cover more grid distance for the same step cost.

## Heuristics

The implementation supports different heuristic functions, which are crucial for guiding the A\* search. You can choose the heuristic when creating a new `Astar` instance. The following heuristics are included:

* **Manhattan Heuristic (`ManhattanHeuristic`)**: Calculates the sum of the absolute differences of the coordinates ($|x_1 - x_2| + |y_1 - y_2|$). This is admissible and consistent for grid movement where only orthogonal steps are allowed (cost 1.0).

* **Diagonal Heuristic (Chebyshev Distance) (`DiagonalHeuristic`)**: Calculates the maximum of the absolute differences of the coordinates ($\max(|x_1 - x_2|, |y_1 - y_2|)$). This is admissible and consistent for grid movement where orthogonal and diagonal steps have a uniform cost (cost 1.0).

* **Euclidean Heuristic (`EuclideanHeuristic`)**: Calculates the straight-line distance between two points ($\sqrt{(x_1 - x_2)^2 + (y_1 - y_2)^2}$). This is admissible and consistent for grid movement where orthogonal steps cost 1.0 and diagonal steps cost $\sqrt{2}$.

## Testing

The project includes table-driven tests to verify the correctness of the `FindPath` function.

To run the tests, clone this repo, navigate to the project directory and run:

```go
go test
```

You can also run specific tests using the -run flag:

```go
go test -run TestFindPath/"path not found"
```

## Benchmarking

Performance benchmarks are included to measure the efficiency of the `FindPath` function on different grid sizes and scenarios.

To run the benchmarks, clone this repo, navigate to the project directory in your terminal and run:

```go
go test -bench=.
```

This will execute all benchmark functions in the package and report metrics like time per operation (ns/op), memory allocations per operation (B/op), and number of allocations per operation (allocs/op).

## Getting Started

1.  Ensure you have Go installed.

2.  In you project directory, run "go get github.com/matteo00gm/go-astar"

3.  Use the `astar.New` function to create an `Astar` instance, providing your grid and desired heuristic.

4.  Call the `astar.FindPath` method with your start and goal coordinates.

Here is an example of how to use this package:

```go
package main

import (
	"fmt"

	"github.com/matteo00gm/go-astar"
)

func main() {

	grid := [][]int{
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 1, 1, 0, 0},
		{0, 1, 1, 1, 1, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	}

	start := astar.Coords{X: 0, Y: 0}
	goal := astar.Coords{X: 9, Y: 3}

	heuristic := &astar.EuclideanHeuristic{}
	astar := astar.New(grid, heuristic)

	path := astar.FindPath(start, goal)
	if path != nil {
		fmt.Println("path found: ", path)
	} else {
		fmt.Println("path not found")
	}
}

// Example Output: path found:  [{0 0} {1 0} {2 0} {2 1} {2 2} {1 2} {0 2} {0 3} {0 4} {1 4} {2 4} {3 4} {4 4} {5 4} {5 3} {5 2} {5 1} {6 1} {7 1} {8 1} {9 2} {9 3}]

```