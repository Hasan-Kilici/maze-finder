package astar

type Point struct {
	X, Y int
}

type Node struct {
	Point
	G, H, F int
	Parent  *Node
}

type PriorityQueue []*Node