# Random Maze Generation and A* Pathfinding Visualization

This project implements a random maze generator and visualizes the solution to the maze using the A* pathfinding algorithm. The maze is generated, rendered as an image, and the path from the start point to the end point is found and highlighted in the image.

## Overview

The main components of this project are:

- **Maze Generation**: A random maze is generated using recursive backtracking.
- **A* Algorithm**: The A* algorithm is utilized to find the shortest path from the start to the end of the maze.
- **Image Rendering**: The maze and its solution are rendered as a PNG image, where:
  - Walls are represented in red.
  - The start point is marked in green.
  - The end point is marked in blue.
  - The path found by the A* algorithm is displayed in yellow.

## File Structure

- **main.go**: The entry point of the application. It generates the maze, saves the image, parses it to extract the start and end points, and finds the path using the A* algorithm.
- **astar/**: Contains the implementation of the A* algorithm, including priority queue management, pathfinding logic, and utility functions for parsing and rendering images.
- **maze/**: Contains the maze generation logic and image rendering functions.

## How to Run

To run the project, ensure you have Go installed on your machine. Then, execute the following commands:

```bash
go run main.go
```

This will generate a maze, save it as `maze.png`, and display the path found by the A* algorithm.
## A* Algorithm

The A* (A-star) algorithm is a popular pathfinding and graph traversal algorithm. It is used for finding the shortest path from a starting node to a goal node while considering the costs of the edges between nodes.
Key Features

- Heuristic Search: A* uses a heuristic to estimate the cost from the current node to the target, which helps in efficiently finding the path.
- Optimality: The algorithm guarantees the shortest path when the heuristic used is admissible, meaning it never overestimates the cost to reach the goal.

## A* Algorithm Pseudocode

- 1.Initialize the open set with the start node.

- 2.While the open set is not empty:
    - Get the node with the lowest f score (where f = g + h).
    - If the current node is the goal, reconstruct the path.
    - Otherwise, move it from the open set to the closed set.
    - For each neighbor of the current node:
        - Calculate the g, h, and f scores.
        - If the neighbor is not in the closed set and not a wall, add it to the open set.

## Maze Generation

The maze is generated using a randomized recursive backtracking algorithm:

- Start with a grid of walls.
- Randomly select a starting point and mark it as a path.
- Recursively visit neighboring cells, carving paths as you go.
- Ensure that there are no isolated sections in the maze.

## References

### A Algorithm*:
- [Wikipedia - A* Search Algorithm](https://en.wikipedia.org/wiki/A*_search_algorithm)
- [Introduction to A* - From Amit’s Thoughts on Pathfinding](http://theory.stanford.edu/~amitp/GameProgramming/AStarComparison.html)

### Maze Generation:
- [Wikipedia - Maze Generation Algorithms](https://en.wikipedia.org/wiki/Maze_generation_algorithm)
