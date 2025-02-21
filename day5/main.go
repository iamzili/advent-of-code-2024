package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// content of pageOrderingRules:
// 29:[13]
// 47:[53 13 61 29]
// 53:[29 13]
// 61:[13 53 29]
// 75:[29 53 47 61 13]
// 97:[13 61 47 29 53 75]]

// e.g updatesToCheck[0] returns:
// [75 47 61 53 29]

// the logic is the following for part1:
// 1. pageOrderingRules[29] = [13] should not be in [75 47 61 53]
// 2. pageOrderingRules[53] = [29 13] should not be in [75 47 61]
// 3. etc ..
// 4. if len(updatesToCheck[0] == 0 we're good

func main() {
	f, _ := os.Open("input")
	// 62|89
	re := regexp.MustCompile(`([0-9]{1,2}\|[0-9]{1,2})`)

	pageOrderingRules := make(map[int][]int)
	updatesToCheck := [][]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if len(scanner.Text()) != 0 {
			if match := re.FindString(line); match != "" {
				substrings := strings.Split(match, "|")

				key, _ := strconv.Atoi(substrings[0])
				val, _ := strconv.Atoi(substrings[1])

				pageOrderingRules[key] = append(pageOrderingRules[key], val)
			} else {
				u := []int{}
				for _, substring := range strings.Split(line, ",") {
					e, _ := strconv.Atoi(substring)
					u = append(u, e)
				}
				updatesToCheck = append(updatesToCheck, u)
			}
		}
	}
	fmt.Println("The sum of the middle page numbers is: ", part1(updatesToCheck, pageOrderingRules))
}

// check if all elements from "subSlice" are present in "slice"
func check(subSlice, slice []int) bool {
	for _, e := range subSlice {
		if slices.Contains(slice, e) {
			return false
		}
	}
	return true
}

func returnMiddle(slice []int) int {
	return slice[len(slice)/2]
}

func part1(updatesToCheck [][]int, pageOrderingRules map[int][]int) int {
	var middlePageNumbers int
	for _, update := range updatesToCheck {
		for i := len(update) - 1; i >= 0; i-- {
			subSlice := update[:i]
			if len(subSlice) == 0 {
				// we finished checking all pages in the update; the order is okay
				middlePageNumbers += returnMiddle(update)
			} else {
				if value, ok := pageOrderingRules[update[i]]; ok {
					if !check(value, subSlice) {
						break
					}
				}
			}
		}
	}
	return middlePageNumbers
}
