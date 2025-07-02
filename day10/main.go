package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input2")

	// grid = map[image.Point]map[image.Point]bool [
	//     		{X: 2, Y: 0}: [
	//				{X: 3, Y: 0}: true,
	//				{X: 2, Y: 1}: true,
	//    	    ],
	//          X: 2, Y: 5}: [
	//          ]
	//			etc ...

	grid := map[image.Point]rune{}
	trailheads := make(map[image.Point]map[image.Point]int)
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r == '0' {
				trailheads[image.Point{x, y}] = make(map[image.Point]int)
			}
			grid[image.Point{x, y}] = r
		}
	}

	delta := []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	candidates := []image.Point{}
	score, rating := 0, 0

	for tr := range trailheads {
		currPos := tr
		for { // check one starting point
			for _, d := range delta { // check every direction at currPos to find a candidate
				if n := currPos.Add(d); int(grid[n])-'0' == ((int(grid[currPos]) - '0') + 1) {
					if int(grid[n])-'0' == 9 {
						trailheads[tr][n] += 1
					}
					candidates = append(candidates, n)
				}
			}
			if candidates == nil {
				break // no candidates found, exit the loop
			} else {
				currPos = candidates[0] // take the first candidate as the new current position
				if len(candidates) != 1 {
					candidates = candidates[1:] // remove the first candidate from the list
				} else {
					candidates = nil
				}
			}
		}
	}
	for _, val := range trailheads {
		score += len(val)
		for _, v := range val {
			rating += v
		}
	}
	fmt.Println("part1: ", score)
	fmt.Println("part2: ", rating)
}
