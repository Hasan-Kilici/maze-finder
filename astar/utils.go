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
	bounds := img.Bounds() // Get the dimensions of the image
	width, height := bounds.Max.X, bounds.Max.Y

	grid := make([][]int, height) // Create a grid with height rows
	for i := range grid {
		grid[i] = make([]int, width) // Each row has a width number of columns
	}

	var start, end Point // Initialize start and end points
	foundStart, foundEnd := false, false // Flags to check if start/end points are found

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
			fmt.Print(grid[y][x], " ") // Print each cell in the grid
		}
		fmt.Println() // New line for each row
	}
}

// PrintPath outputs the found path to the console
func PrintPath(path []Point) {
	for _, p := range path {
		fmt.Printf("(%d, %d) -> ", p.X, p.Y) // Print each point in the path
	}
	fmt.Println("Reached the goal!") // Indicate the goal has been reached
}

// DrawPathOnImage overlays the found path on the original image and saves it as a new file
func DrawPathOnImage(img image.Image, path []Point) {
	bounds := img.Bounds() // Get the bounds of the original image
	newImg := image.NewRGBA(bounds) // Create a new RGBA image to draw on

	// Copy original image to the new image
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newImg.Set(x, y, img.At(x, y)) // Set each pixel in the new image
		}
	}

	yellow := color.RGBA{255, 255, 0, 255} // Define the color yellow for the path
	for _, p := range path {
		newImg.Set(p.X, p.Y, yellow) // Set the path points to yellow in the new image
	}

	outFile, err := os.Create("./output/output.png") // Create output file
	if err != nil {
		fmt.Println("Error while saving the new image:", err)
		return
	}
	defer outFile.Close() // Ensure the file is closed after writing

	png.Encode(outFile, newImg) // Encode and save the new image
	fmt.Println("Path successfully drawn and saved as 'output.png'.") // Confirmation message
}
