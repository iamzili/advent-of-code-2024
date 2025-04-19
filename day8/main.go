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
					calcUpperAntidote(a, b, false, grid, part1)  // part1
					calcBottomAntidote(a, b, false, grid, part1) // part1
					// antinodes appear on every antenna in part2
					part2[a] = true
					part2[b] = true
					calcUpperAntidote(a, b, true, grid, part2)  // part2
					calcBottomAntidote(a, b, true, grid, part2) // part2

				}
			}
		}
	}
	fmt.Println("Part1: ", len(part1))
	fmt.Println("Part2: ", len(part2))
}

// To calculate the coordinates of the upper antinode for a pair of antennas, use the following formula:
// (firstAntenna - secondAntenna) + firstAntenna
func calcUpperAntidote(a, b image.Point, p2 bool, grid map[image.Point]rune, antidotes map[image.Point]bool) {
	upperAntidote := a.Sub(b).Add(a)
	if grid[upperAntidote] != 0 {
		antidotes[upperAntidote] = true
	} else {
		p2 = false
	}
	if p2 {
		calcUpperAntidote(upperAntidote, a, true, grid, antidotes)
	}
}

// To calculate the coordinates of the bottom antinode for a pair of antennas, use the following formula:
// (secondAntenna - firstAntenna) + secondAntenna
func calcBottomAntidote(a, b image.Point, p2 bool, grid map[image.Point]rune, antidotes map[image.Point]bool) {
	bottomAntidote := b.Sub(a).Add(b)
	if grid[bottomAntidote] != 0 {
		antidotes[bottomAntidote] = true
	} else {
		p2 = false
	}
	if p2 {
		calcBottomAntidote(b, bottomAntidote, true, grid, antidotes)
	}
}
