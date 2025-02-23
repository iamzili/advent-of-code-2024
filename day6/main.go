package main

import (
	"fmt"
	"os"
	"slices"
)

const obstacle, guard, eol byte = 35, 94, 10

func main() {
	input, _ := os.ReadFile("input")
	size := slices.Index(input, eol)
	currPos := slices.Index(input, guard)
	//startPos := currPos
	currDirection := "up"
	// add the initial position
	visited := map[int]int{
		currPos: 1,
	}
	move := true

	for move {
		switch {
		case currDirection == "up":
			currPos -= size + 1
			if !indexExists(input, currPos) { // We found the way out
				move = false
				break
			}
			if isObstacle(input, currPos) { // start moving right
				currDirection = "right"
				currPos += size + 1 // Go back to the previous position
			} else {
				visited[currPos] += 1
			}
		case currDirection == "down":
			currPos += size + 1
			if !indexExists(input, currPos) {
				fmt.Println("We found the way out!")
				move = false
				break
			}
			if isObstacle(input, currPos) { // start moving left
				currDirection = "left"
				currPos -= size + 1 //Go back to the previous position
			} else {
				visited[currPos] += 1
			}
		case currDirection == "right":
			currPos += 1
			if isObstacle(input, currPos) { // start moving down
				currDirection = "down"
				currPos -= 1 // Go back to the previous position
			} else if input[currPos] == eol { // We found the way out
				move = false
				break
			} else {
				visited[currPos] += 1
			}
		case currDirection == "left":
			currPos -= 1
			// We found the way out
			if !indexExists(input, currPos) || input[currPos] == eol {
				move = false
				break
			}
			if isObstacle(input, currPos) { // start moving up
				currDirection = "up"
				currPos += 1 //Go back to the previous position
			} else {
				visited[currPos] += 1
			}
		}
	}
	fmt.Println("distinct positions visited by the guard is ", len(visited))
}

func indexExists(slice []byte, index int) bool {
	return index >= 0 && index < len(slice)
}

func isObstacle(slice []byte, index int) bool {
	return slice[index] == obstacle
}
