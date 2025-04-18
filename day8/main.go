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
	uniqueAntidotes := map[image.Point]bool{}
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
					// To calculate the coordinates of the upper antinode for a pair of antennas, use the following formula:
					// (firstAntenna - secondAntenna) + firstAntenna
					upperAntidote := a.Sub(b).Add(a)
					if grid[upperAntidote] != 0 {
						uniqueAntidotes[upperAntidote] = true
					}
					// To calculate the coordinates of the bottom antinode for a pair of antennas, use the following formula:
					// (secondAntenna - firstAntenna) + secondAntenna
					bottomAntidote := b.Sub(a).Add(b)
					if grid[bottomAntidote] != 0 {
						uniqueAntidotes[bottomAntidote] = true
					}
				}
			}
		}
	}
	fmt.Println("Part1: ", len(uniqueAntidotes))
}
