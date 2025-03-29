package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

func calcDay1(testvalue int, numbers []string) int {

	symbols := []string{"+", "-"}
	equations := calcCartesianProduct(len(numbers)-1, symbols)

	for _, equation := range equations {
		if calcEquation(testvalue, equation, numbers) {
			return testvalue
		}
	}
	return 0
}

func calcDay2(testvalue int, numbers []string) int {

	symbols := []string{"+", "-", "||"}
	equations := calcCartesianProduct(len(numbers)-1, symbols)

	for _, equation := range equations {
		if calcEquation(testvalue, equation, numbers) {
			return testvalue
		}
	}
	return 0
}

// When generating all combinations of "+" and "-" with a fixed length n
// you're essentially computing the n-fold Cartesian product of the set {+, -} with itself
// For example, if A = {+, -}, the Cartesian product A × A would give:
// {(+,+), (+,-), (-,+), (-,-)}
func calcCartesianProduct(length int, symbols []string) [][]int {
	// Create dimensions array for Cartesian product
	dims := make([]int, length)
	for i := range dims {
		dims[i] = len(symbols)
	}
	// generate all combinations using Cartesian product
	list := combin.Cartesian(dims)
	return list
}

func calcEquation(testvalue int, equation []int, numbers []string) bool {
	result, _ := strconv.Atoi(numbers[0])
	// start looping from the second index
	for i, s := range numbers[1:] {
		num, _ := strconv.Atoi(s)
		if equation[i] == 0 {
			result += num
		} else if equation[i] == 1 {
			result *= num
		} else if equation[i] == 2 {
			result, _ = strconv.Atoi(strconv.Itoa(result) + s)
		}
	}
	if testvalue == result {
		return true
	} else {
		return false
	}
}

func main() {
	input, _ := os.ReadFile("input2")
	//fmt.Println(string(input))
	var day1, day2 int

	lines := strings.Split(string(input), "\n")
	for _, s := range lines {
		v := strings.Split(s, ":")
		x, _ := strconv.Atoi(v[0])
		day1 += calcDay1(x, strings.Fields(v[1]))
		day2 += calcDay2(x, strings.Fields(v[1]))
	}

	fmt.Println("Day1: ", day1)
	fmt.Println("Day2: ", day2)
}
