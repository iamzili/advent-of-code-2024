package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	fmt.Println("day9")

	input, _ := os.ReadFile("input")
	block := []rune{} // For input2, this slice after the first for loop should contain: "00...111...2...333.44.5555.6666.777.888899"
	checksum := 0

	if cap(input) <= len(input) {
		log.Fatalln("the capacity of the slice is equal to or less than its length")
	}

	for i := 0; i < (len(input)+1)/2; i++ {
		pair := input[i*2 : 2+(i*2)]
		layout, _ := strconv.Atoi(string(pair[0]))
		fspace, _ := strconv.Atoi(string(pair[1]))
		for range layout {
			block = append(block, rune('0'+i))
		}
		for range fspace {
			block = append(block, '.')
		}
	}

	lastIndex := len(block) - 1
	for firstIndex := 0; firstIndex < len(block); firstIndex++ {
		if block[firstIndex] != '.' {
			// Convert int to rune -> int(rune), getting integer for specific ASCII character, just subtract it.
			checksum += firstIndex * (int(block[firstIndex]) - '0')
		} else if block[firstIndex] == '.' {
			for tempLastIndex, b := range slices.Backward(block) {
				if b == '.' {
					continue
				} else { // found a file block
					lastIndex = tempLastIndex
					checksum += firstIndex * (int(block[lastIndex]) - '0')
					block = block[:lastIndex]
					break
				}
			}
		}
		if firstIndex == lastIndex {
			break
		}
	}
	fmt.Printf("\n")
	fmt.Println("part1: ", checksum)
}
