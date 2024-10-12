package maze

import (
    "math/rand"
)

// Cell represents a single cell in the maze
type Cell struct {
    IsWall bool // Indicates if the cell is a wall
}

// Directions for moving in the maze (right, down, left, up)
var directions = [][2]int{
    {0, 1},   // Move right
    {1, 0},   // Move down
    {0, -1},  // Move left
    {-1, 0},  // Move up
}

// GenerateMaze creates a new maze of the specified width and height
func GenerateMaze(width, height int) [][]Cell {
    maze := make([][]Cell, height)
    for y := range maze {
        maze[y] = make([]Cell, width)
    }

    // Fill the maze with walls
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            maze[y][x] = Cell{IsWall: true} // Set all cells as walls initially
        }
    }

    // Randomly select a starting point within the maze
    startX := rand.Intn((width-2)/2)*2 + 1
    startY := rand.Intn((height-2)/2)*2 + 1

    maze[startY][startX].IsWall = false // Mark the starting cell as open

    // Start the maze generation process
    visit(maze, startX, startY)

    // Ensure the start and end points are open
    maze[startY][startX] = Cell{IsWall: false}
    maze[height-2][width-2] = Cell{IsWall: false} // Set the end cell to open

    return maze
}

// visit recursively visits cells in the maze to create paths
func visit(maze [][]Cell, x, y int) {
    // Shuffle directions to create a random maze layout
    rand.Shuffle(len(directions), func(i, j int) {
        directions[i], directions[j] = directions[j], directions[i]
    })

    // Explore each direction
    for _, direction := range directions {
        nx, ny := x+direction[0]*2, y+direction[1]*2 // Calculate the next cell

        // Check if the next cell is within bounds
        if nx >= 0 && ny >= 0 && nx < len(maze[0]) && ny < len(maze) {
            if maze[ny][nx].IsWall { // If the next cell is a wall
                // Carve a path to the next cell
                maze[y+direction[1]][x+direction[0]] = Cell{IsWall: false}
                maze[ny][nx] = Cell{IsWall: false}

                visit(maze, nx, ny)
            }
        }
    }
}
