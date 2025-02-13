package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func StringToIntSlice(str string) []int {
	var intSlice []int
	strSlice := strings.Split(str, " ")

	for _, e := range strSlice {
		if e, err := strconv.Atoi(e); err == nil {
			intSlice = append(intSlice, e)
		}
	}
	return intSlice
}

func main() {
	f, _ := os.Open("input")
	var intSlice []int
	var numSafeReps int

	var shouldDesc, shouldAsc bool = false, false

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		intSlice = StringToIntSlice(scanner.Text())
		shouldDesc, shouldAsc = false, false

		if slices.IsSortedFunc(intSlice, func(a, b int) int {
			// this decides initially the direction
			if !shouldDesc && !shouldAsc {
				if a < b {
					shouldDesc = true
				} else if a > b {
					shouldAsc = true
				}
			}
			if a < b && !shouldAsc {
				if diff := b - a; diff >= 1 && diff <= 3 {
					return 1
				}
				return -1
			} else if a > b && !shouldDesc {
				if diff := a - b; diff >= 1 && diff <= 3 {
					return 1
				}
				return -1
			}
			return -1
		}) {
			// Levels are either all increasing or all decreasing and levels differ by at least 1 and at most 3
			numSafeReps++
		}
	}
	fmt.Println("The number of reports which are safe is ", numSafeReps)
}
