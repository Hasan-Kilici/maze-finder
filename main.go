package main

import (
    "pathfinder/maze"
    "pathfinder/astar"
    "fmt"
    "image/png"
    "os"
    "log"
    "image"
)

func main() {
    width, height := 500, 200

    img := maze.RenderMaze(maze.GenerateMaze(width, height))

    if err := maze.SaveImage("maze.png", img); err != nil {
        log.Fatal(err)
    }

    file, err := os.Open("maze.png")
    if err != nil {
        log.Fatalf("Error while opening image: %v", err)
    }
    defer file.Close()

    rgbaImg, err := png.Decode(file)
    if err != nil {
        log.Fatalf("Error while decoding PNG: %v", err)
    }

    grid, start, end, err := astar.ParseImage(rgbaImg.(*image.RGBA))
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Initial Grid:")
    astar.PrintGrid(grid)

    if path := astar.FindPath(grid, start, end); path != nil {
        fmt.Println("Path found:")
        astar.PrintPath(path)
        astar.DrawPathOnImage(rgbaImg.(*image.RGBA), path)
    } else {
        fmt.Println("No path found.")
    }
}
