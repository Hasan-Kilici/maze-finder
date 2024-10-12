package astar

import (
	"container/heap"
	"math"
)

const (
	EMPTY = 0
	WALL  = 1
	START = 2
	END   = 3
)

// Directions for movement in the grid (up, right, down, left)
var directions = []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

// Calculates the Manhattan distance between two points
func manhattanDistance(p1, p2 Point) int {
	return int(math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y)))
}

// FindPath implements the A* pathfinding algorithm
func FindPath(grid [][]int, start, end Point) []Point {
	openSet := &PriorityQueue{}
	heap.Init(openSet)
	closedSet := make(map[Point]bool)

	// Create the start node
	startNode := &Node{Point: start, G: 0, H: manhattanDistance(start, end)}
	startNode.F = startNode.G + startNode.H // Calculate F score
	heap.Push(openSet, startNode)

	// Loop until there are nodes in the open set
	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*Node)

		// Check if we've reached the end point
		if current.Point == end {
			var path []Point
			// Construct the path by backtracking from the end node to the start node
			for node := current; node != nil; node = node.Parent {
				path = append([]Point{node.Point}, path...)
			}
			return path
		}

		closedSet[current.Point] = true // Mark the current node as closed

		// Explore neighbors (up, right, down, left)
		for _, dir := range directions {
			neighbor := Point{current.X + dir.X, current.Y + dir.Y} // Calculate the neighbor point

			// Check if the neighbor is valid for exploration
			if isValid(neighbor, grid, closedSet) {
				g := current.G + 1 // Calculate the cost to reach the neighbor
				h := manhattanDistance(neighbor, end) // Heuristic cost to reach the end
				f := g + h // Total cost for the neighbor

				// Create a new node for the neighbor
				neighborNode := &Node{Point: neighbor, G: g, H: h, F: f, Parent: current}
				heap.Push(openSet, neighborNode) // Push neighbor node to open set
			}
		}
	}

	return nil // Return nil if no path is found
}

// Checks if a point is valid for exploration
func isValid(p Point, grid [][]int, closedSet map[Point]bool) bool {
	return p.X >= 0 && p.X < len(grid[0]) && p.Y >= 0 && p.Y < len(grid) && grid[p.Y][p.X] != WALL && !closedSet[p]
}
