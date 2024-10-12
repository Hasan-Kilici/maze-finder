package maze

import (
    "image"
    "image/color"
    "image/png"
    "os"
)

const (
    wallColor  = 0xFF0000
    startColor = 0x00FF00
    endColor   = 0x0000FF
    pathColor  = 0xFFFFFF
)

// RenderMaze creates an image representation of the maze
func RenderMaze(maze [][]Cell) *image.RGBA {
    width := len(maze[0])
    height := len(maze)

    img := image.NewRGBA(image.Rect(0, 0, width, height))

    wallColor := color.RGBA{255, 0, 0, 255}
    pathColor := color.RGBA{0, 0, 0, 255}
    startColor := color.RGBA{0, 255, 0, 255}
    endColor := color.RGBA{0, 0, 255, 255}

    // Iterate through the maze to set pixel colors
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            if maze[y][x].IsWall {
                img.Set(x, y, wallColor)
            } else {
                img.Set(x, y, pathColor)
            }
        }
    }

    startX := 1
    startY := 1
    img.Set(startX, startY, startColor)

    endX := width - 2
    endY := height - 2
    img.Set(endX, endY, endColor)

    return img 
}

// SaveImage saves the generated image to a file
func SaveImage(filename string, img *image.RGBA) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    return png.Encode(file, img)
}
