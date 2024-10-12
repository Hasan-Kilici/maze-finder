package maze

import (
    "image"
    "image/color"
    "image/png"
    "os"
)

const (
    wallColor  = 0xFF0000 // Red color for walls
    startColor = 0x00FF00 // Green color for the start point
    endColor   = 0x0000FF // Blue color for the end point
    pathColor  = 0xFFFFFF // White color for paths
)

// RenderMaze creates an image representation of the maze
func RenderMaze(maze [][]Cell) *image.RGBA {
    width := len(maze[0]) // Get the width of the maze
    height := len(maze)    // Get the height of the maze

    // Create a new RGBA image with the dimensions of the maze
    img := image.NewRGBA(image.Rect(0, 0, width, height))

    // Define colors using color.RGBA
    wallColor := color.RGBA{255, 0, 0, 255}   // Red for walls
    pathColor := color.RGBA{0, 0, 0, 255}      // Black for paths
    startColor := color.RGBA{0, 255, 0, 255}   // Green for the start point
    endColor := color.RGBA{0, 0, 255, 255}     // Blue for the end point

    // Iterate through the maze to set pixel colors
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            if maze[y][x].IsWall {
                img.Set(x, y, wallColor) // Set wall color
            } else {
                img.Set(x, y, pathColor) // Set path color
            }
        }
    }

    // Set the start point color
    startX := 1
    startY := 1
    img.Set(startX, startY, startColor)

    // Set the end point color
    endX := width - 2
    endY := height - 2
    img.Set(endX, endY, endColor)

    return img // Return the generated image
}

// SaveImage saves the generated image to a file
func SaveImage(filename string, img *image.RGBA) error {
    file, err := os.Create(filename) // Create a file for output
    if err != nil {
        return err // Return any error encountered during file creation
    }
    defer file.Close() // Ensure the file is closed after function execution

    return png.Encode(file, img) // Encode the image in PNG format and save it
}
