package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var blinks = 25

func main() {
	input, _ := os.ReadFile("input2")
	stones := strings.Split(string(input), " ")
	currBlinks := 0

	fmt.Println("part1: ", part1(stones, currBlinks))
}

func part1(stones []string, currBlinks int) int {
	tmpStones := []string{}
	if currBlinks != blinks {
		currBlinks++
		for _, s := range stones {
			if s == "0" {
				tmpStones = append(tmpStones, "1")
			} else if len(s)%2 == 0 {
				stone := strings.SplitAfter(s, s[:len(s)/2])
				s1, _ := strconv.Atoi(stone[0])
				s2, _ := strconv.Atoi(stone[1])
				tmpStones = append(tmpStones, strconv.Itoa(s1), strconv.Itoa(s2))
			} else {
				i, _ := strconv.Atoi(s)
				i *= 2024
				tmpStones = append(tmpStones, strconv.Itoa(i))
			}
		}
		return part1(tmpStones, currBlinks)
	}
	return len(stones)
}
