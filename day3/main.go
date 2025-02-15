package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(data []byte, re regexp.Regexp) {
	var allInst int
	matches := re.FindAllString(string(data), -1)

	for _, e := range matches {
		f := func(c rune) bool {
			return !unicode.IsNumber(c)
		}
		inst := strings.FieldsFunc(e, f)
		x, _ := strconv.Atoi(inst[0])
		y, _ := strconv.Atoi(inst[1])
		allInst += x * y
	}
	fmt.Println("Part 1 - Sum of all uncorrupted mul instructions is ", allInst)
}

func part2(data []byte, re regexp.Regexp) {
	var allInst int
	var do bool = true
	matches := re.FindAllString(string(data), -1)

	for _, e := range matches {
		if e == "do()" {
			do = true
			continue
		} else if e == "don't()" {
			do = false
			continue
		}
		if do {
			f := func(c rune) bool {
				return !unicode.IsNumber(c)
			}
			inst := strings.FieldsFunc(e, f)
			x, _ := strconv.Atoi(inst[0])
			y, _ := strconv.Atoi(inst[1])
			allInst += x * y
		}
	}
	fmt.Println("Part2 - Sum of all uncorrupted mul instructions is ", allInst)
}

func main() {
	data, err := os.ReadFile("input")
	check(err)

	var re1 = regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	part1(data, *re1)

	var re2 = regexp.MustCompile(`(don't\()\)|(do\(\))|(mul\([0-9]{1,3},[0-9]{1,3}\))`)
	part2(data, *re2)
}
