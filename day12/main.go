package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

var delta = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} // up, right, down, left

func main() {
	input, _ := os.ReadFile("input4")

	grid := make(map[image.Point]rune)
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			grid[image.Point{x, y}] = r
		}
	}
	fmt.Println("part1:", part1(grid))
}

func part1(grid map[image.Point]rune) int {
	seen := map[image.Point]int{}
	totalPrice := 0

	for plot := range grid {
		perimeter, area := getPerimeterArea(grid, seen, plot)
		totalPrice += perimeter * area
	}
	return totalPrice
}

func getPerimeterArea(grid map[image.Point]rune, seen map[image.Point]int, currPlot image.Point) (int, int) {
	perimeter, area := 0, 0

	if seen[currPlot] == 0 { // if the current plot has not been visited yet
		area++             // increase the area of the current plot
		seen[currPlot] = 1 // mark the current plot as visited
	} else {
		return 0, 0 // if the current plot has been visited
	}
	for _, d := range delta {
		next := currPlot.Add(d)
		if nextPlotPlant, ok := grid[next]; !ok { // if next plot is out of bounds, increase the perimeter
			perimeter++
		} else if nextPlotPlant != grid[currPlot] { // if the next plot has not the same region as the current plot increase the perimeter
			perimeter++
		} else if nextPlotPlant == grid[currPlot] && seen[next] == 0 { // if the next plot has the same region as the current plot
			p, a := getPerimeterArea(grid, seen, next)
			perimeter += p
			area += a
		}
	}
	return perimeter, area
}
