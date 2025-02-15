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
	fmt.Println("Sum of all uncorrupted mul instructions is ", allInst)
}

func main() {
	dat, err := os.ReadFile("input")
	check(err)

	var re = regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	part1(dat, *re)
}
