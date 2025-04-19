package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input2")

	grid := map[image.Point]rune{}
	antennas := make(map[rune][]image.Point)
	part1 := map[image.Point]bool{}
	part2 := map[image.Point]bool{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			grid[image.Point{x, y}] = r
			if r != '.' {
				antennas[r] = append(antennas[r], image.Point{x, y})
			}
		}
	}

	for freq := range antennas {
		if len(antennas[freq]) >= 2 { // antinode occurs at any point that is perfectly in line with two antennas of the same frequency
			// By using the two for loops, we obtain all possible subsets (k-subsets) containing exactly 2 elements from the antennas[freq] set
			for i, a := range antennas[freq] {
				for _, b := range antennas[freq][i+1:] {
					calcUpperAntinode(a, b, false, grid, part1)  // part1
					calcBottomAntinode(a, b, false, grid, part1) // part1
					// antinodes appear on every antenna in part2
					part2[a] = true
					part2[b] = true
					calcUpperAntinode(a, b, true, grid, part2)  // part2
					calcBottomAntinode(a, b, true, grid, part2) // part2

				}
			}
		}
	}
	fmt.Println("Part1: ", len(part1))
	fmt.Println("Part2: ", len(part2))
}

// To calculate the coordinates of the upper antinode for a pair of antennas, use the following formula:
// (firstAntenna - secondAntenna) + firstAntenna
func calcUpperAntinode(a, b image.Point, p2 bool, grid map[image.Point]rune, antinodes map[image.Point]bool) {
	upperAntinode := a.Sub(b).Add(a)
	if grid[upperAntinode] != 0 {
		antinodes[upperAntinode] = true
	} else {
		p2 = false
	}
	if p2 {
		calcUpperAntinode(upperAntinode, a, true, grid, antinodes)
	}
}

// To calculate the coordinates of the bottom antinode for a pair of antennas, use the following formula:
// (secondAntenna - firstAntenna) + secondAntenna
func calcBottomAntinode(a, b image.Point, p2 bool, grid map[image.Point]rune, antinodes map[image.Point]bool) {
	bottomAntinode := b.Sub(a).Add(b)
	if grid[bottomAntinode] != 0 {
		antinodes[bottomAntinode] = true
	} else {
		p2 = false
	}
	if p2 {
		calcBottomAntinode(b, bottomAntinode, true, grid, antinodes)
	}
}
