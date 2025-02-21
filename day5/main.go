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

// the logic is the following for isOrdered func:
// 1. pageOrderingRules[29] = [13] should not be in [75 47 61 53]
// 2. pageOrderingRules[53] = [29 13] should not be in [75 47 61]
// 3. etc ..

func main() {
	f, _ := os.Open("input")
	// 62|89
	re := regexp.MustCompile(`([0-9]{1,2}\|[0-9]{1,2})`)

	pageOrderingRules := make(map[int][]int)
	updatesToCheck := [][]int{}
	var part1, part2 int

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
				update := []int{}
				for _, substring := range strings.Split(line, ",") {
					pageNum, _ := strconv.Atoi(substring)
					update = append(update, pageNum)
				}
				updatesToCheck = append(updatesToCheck, update)
			}
		}
	}

	for _, update := range updatesToCheck {
		if isOrdered(update, pageOrderingRules) {
			part1 += returnMiddle(update)
		} else {
			update = fixOrder(update, pageOrderingRules)
			part2 += returnMiddle(update)
		}
	}

	fmt.Println("The sum of the middle page numbers for ordered updates is: ", part1)
	fmt.Println("The sum of the middle page numbers after correctly ordering is: ", part2)
}

// check if all elements from "subSlice" are present in "slice"
func check(rule, subSlice []int) (bool, []int) {
	containsMapping := make([]int, len(subSlice))
	var isPresent bool
	for i, e := range subSlice {
		if slices.Contains(rule, e) {
			containsMapping[i] = 1
			isPresent = true
		} else {
			containsMapping[i] = 0
		}
	}
	return isPresent, containsMapping
}

func returnMiddle(slice []int) int {
	return slice[len(slice)/2]
}

func isOrdered(update []int, pageOrderingRules map[int][]int) bool {
	for i := len(update) - 1; i >= 0; i-- {
		subSlice := update[:i]
		if rule, ok := pageOrderingRules[update[i]]; ok {
			// incorrectly-ordered
			if ok, _ := check(rule, subSlice); ok {
				return false
			}
		}
	}
	return true
}

func fixOrder(update []int, pageOrderingRules map[int][]int) []int {
	for i := len(update) - 1; i >= 0; i-- {
		subSlice := update[:i]
		if rule, ok := pageOrderingRules[update[i]]; ok {
			// incorrectly-ordered
			if ok, containsMapping := check(rule, subSlice); ok {
				fixOrder(fixOneStep(update, containsMapping, i), pageOrderingRules)
			}
		}
	}
	return update
}

func fixOneStep(update, containsMapping []int, index int) []int {
	for i, e := range containsMapping {
		if e == 1 {
			val := update[index]
			// move a value
			update = slices.Delete(update, index, index+1)
			update = slices.Insert(update, i, val)
			break
		}
	}
	return update
}
