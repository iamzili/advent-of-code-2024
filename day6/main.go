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
	currGuardPos := slices.Index(input, guard)
	currDirection := "up"
	visited := []int{currGuardPos} // add the initial position
	move := true

	for move {
		switch {
		case currDirection == "up":
			currGuardPos -= size + 1
			if !indexExists(input, currGuardPos) { // We found the way out
				move = false
				break
			}
			if isObstacle(input, currGuardPos) { // start moving right
				currDirection = "right"
				currGuardPos += size + 1 // Go back to the previous position
			} else {
				visited = append(visited, currGuardPos)
			}
		case currDirection == "down":
			currGuardPos += size + 1
			if !indexExists(input, currGuardPos) {
				fmt.Println("We found the way out!")
				move = false
				break
			}
			if isObstacle(input, currGuardPos) { // start moving left
				currDirection = "left"
				currGuardPos -= size + 1 //Go back to the previous position
			} else {
				visited = append(visited, currGuardPos)
			}
		case currDirection == "right":
			currGuardPos += 1
			if isObstacle(input, currGuardPos) { // start moving down
				currDirection = "down"
				currGuardPos -= 1 // Go back to the previous position
			} else if input[currGuardPos] == eol { // We found the way out
				move = false
				break
			} else {
				visited = append(visited, currGuardPos)
			}
		case currDirection == "left":
			currGuardPos -= 1
			// We found the way out
			if !indexExists(input, currGuardPos) || input[currGuardPos] == eol {
				move = false
				break
			}
			if isObstacle(input, currGuardPos) { // start moving up
				currDirection = "up"
				currGuardPos += 1 //Go back to the previous position
			} else {
				visited = append(visited, currGuardPos)
			}
		}
	}
	slices.Sort(visited)
	visited = slices.Compact(visited) // remove duplicated postions
	fmt.Println("distinct positions visited by the guard is ", len(visited))
}

func indexExists(slice []byte, index int) bool {
	return index >= 0 && index < len(slice)
}

func isObstacle(slice []byte, index int) bool {
	return slice[index] == obstacle
}
