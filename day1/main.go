package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

// Removes slice element at index(s) and returns new slice
func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func main() {

	var left []int
	var right []int
	var dist int

	f, _ := os.Open("input")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")

		w1, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}
		w2, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}

		left = append(left, w1)
		right = append(right, w2)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	slices.Sort(left)
	slices.Sort(right)

	part1 := func() int {
		if len(left) == len(right) {
			for i := range left {
				dist += abs(left[i] - right[i])
			}
		}
		return dist
	}

	part2 := func() int {
		var count, simScore int
		for _, e := range left {
			for {
				n, found := slices.BinarySearch(right, e)
				if found {
					right = remove(right, n)
					count++
				} else {
					simScore += e * count
					count = 0
					break
				}
			}
		}
		return simScore
	}

	fmt.Println("total distance between the left list and the right list is: ", part1())
	fmt.Println("similarity score is: ", part2())
}
