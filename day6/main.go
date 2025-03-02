package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

// Credit: https://github.com/mnml/aoc/tree/main/2024

func main() {
	input, _ := os.ReadFile("input2")

	// grid = map[image.Point]int32 [
	//     		{X: 0, Y: 0}: 46,
	//     		{X: 1, Y: 0}: 46,
	//			etc ...
	// start = image.Point {X: 4, Y: 6}

	grid, start := map[image.Point]rune{}, image.Point{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r == '^' {
				start = image.Point{x, y}
			}
			grid[image.Point{x, y}] = r
		}
	}

	// delta = []image.Point len: 4, cap: 4, [
	//    {X: 0, Y: -1},     up,    delta[0] aka 0
	//    {X: 1, Y: 0},      right, delta[1] aka 1
	//    {X: 0, Y: 1},      down,  delta[2] aka 2
	//    {X: -1, Y: 0},     left,  delta[3] aka 3
	//]

	patrol := func(o image.Point) map[image.Point]int {
		delta := []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
		// d contains the direction
		p, d, seen := start, 0, map[image.Point]int{}
		for {
			if _, ok := grid[p]; !ok {
				// found the exit
				return seen
			} else if 1<<d&seen[p] != 0 { // 1. seen[p] ensures that p needs to be in the map, if not then x & 0 = 0
				// we found a loop
				return nil
			}
			seen[p] |= 1 << d                                   // add a step to the "seen" map
			if n := p.Add(delta[d]); grid[n] == '#' || n == o { // n == o needed for exit
				d = (d + 1) % len(delta) // set new direction, e.g. if the curr direction is 0 which is up then: 1 % 4 = 1 aka right is going to be next direction
			} else {
				p = n
			}
		}
	}

	part1, part2 := patrol(image.Point{-1, -1}), 0
	for p := range part1 {
		if patrol(p) == nil {
			part2++
		}
	}
	fmt.Println(len(part1))
	fmt.Println(part2)
}
