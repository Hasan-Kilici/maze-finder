package astar

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// ParseImage converts an image into a grid representation and identifies the start and end points
func ParseImage(img image.Image) ([][]int, Point, Point, error) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	var start, end Point
	foundStart, foundEnd := false, false

	// Loop through each pixel in the image
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA() // Get the RGBA values of the pixel
			r, g, b = r>>8, g>>8, b>>8 // Scale down the color values

			switch {
			case r == 255 && g == 0 && b == 0: // Red represents a wall
				grid[y][x] = WALL
			case r == 0 && g == 255 && b == 0: // Green represents the start point
				grid[y][x] = START
				start = Point{X: x, Y: y} // Set start point
				foundStart = true
			case r == 0 && g == 0 && b == 255: // Blue represents the end point
				grid[y][x] = END
				end = Point{X: x, Y: y} // Set end point
				foundEnd = true
			default:
				grid[y][x] = EMPTY // Any other color is an empty space
			}
		}
	}

	// Check if start and end points were found
	if !foundStart || !foundEnd {
		return nil, Point{}, Point{}, fmt.Errorf("Start (green) or End (blue) point not found in the image")
	}

	return grid, start, end, nil // Return the grid and identified points
}

// PrintGrid outputs the grid to the console
func PrintGrid(grid [][]int) {
	for y := range grid {
		for x := range grid[y] {
			fmt.Print(grid[y][x], " ")
		}
		fmt.Println()
	}
}

func PrintPath(path []Point) {
	for _, p := range path {
		fmt.Printf("(%d, %d) -> ", p.X, p.Y)
	}
	fmt.Println("Reached the goal!")
}

// DrawPathOnImage overlays the found path on the original image and saves it as a new file
func DrawPathOnImage(img image.Image, path []Point) {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newImg.Set(x, y, img.At(x, y))
		}
	}

	yellow := color.RGBA{255, 255, 0, 255}
	for _, p := range path {
		newImg.Set(p.X, p.Y, yellow)
	}

	outFile, err := os.Create("./output/output.png")
	if err != nil {
		fmt.Println("Error while saving the new image:", err)
		return
	}
	defer outFile.Close()

	png.Encode(outFile, newImg)
	fmt.Println("Path successfully drawn and saved as 'output.png'.")
}
