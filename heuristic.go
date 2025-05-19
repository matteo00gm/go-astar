package astar

import "math"

type Heuristic interface {
	estimate(pos_a Coords, pos_b Coords) float64
}

type ManhattanHeuristic struct{}

func (m *ManhattanHeuristic) estimate(pos_a, pos_b Coords) float64 {
	return math.Abs(float64(pos_a.X-pos_b.X)) + math.Abs(float64(pos_a.Y-pos_b.Y))
}

type DiagonalHeuristic struct{}

func (m *DiagonalHeuristic) estimate(pos_a, pos_b Coords) float64 {
	return max(math.Abs(float64(pos_a.X-pos_b.X)), math.Abs(float64(pos_a.Y-pos_b.Y)))
}

type EuclideanHeuristic struct{}

func (m *EuclideanHeuristic) estimate(pos_a, pos_b Coords) float64 {
	return math.Sqrt(float64((pos_a.X-pos_b.X)*(pos_a.X-pos_b.X) + (pos_a.Y-pos_b.Y)*(pos_a.Y-pos_b.Y)))
}
