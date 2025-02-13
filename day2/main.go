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

func part1(slice []int) bool {
	var shouldDesc, shouldAsc bool = false, false

	if slices.IsSortedFunc(slice, func(a, b int) int {
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
		return true
	} else {
		// not safe report
		return false
	}
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	f, _ := os.Open("input")
	var numSafeReps, numSafeRepsPD int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		slice := StringToIntSlice(scanner.Text())
		lenght := len(slice)
		if part1(slice) {
			numSafeReps++
		} else {
			newSlice := make([]int, len(slice))
			for i := lenght - 1; i >= 0; i-- {
				copy(newSlice, slice)
				if part1(removeIndex(newSlice, i)) {
					numSafeRepsPD++
					break
				}
			}
		}
	}
	fmt.Println("The number of reports which are safe is ", numSafeReps)
	fmt.Println("The number of reports which are safe using the Problem Dampener is: ", numSafeReps+numSafeRepsPD)
}
