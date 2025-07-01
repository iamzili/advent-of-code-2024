package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("input")
	block := []rune{} // For input2, this slice after the first for loop should contain: "00...111...2...333.44.5555.6666.777.888899"
	tmpBlock := []rune{}
	block2 := [][]rune{} // for part2 I'm using a different approach, i.e. I keep the file blocks in a 2D slice

	if cap(input) <= len(input) {
		log.Fatalln("the capacity of the slice is equal to or less than its length")
	}

	for i := 0; i < (len(input)+1)/2; i++ {
		pair := input[i*2 : 2+(i*2)]
		layout, _ := strconv.Atoi(string(pair[0]))
		fspace, _ := strconv.Atoi(string(pair[1]))

		for range layout {
			tmpBlock = append(tmpBlock, rune('0'+i))
		}
		block = slices.Concat(block, tmpBlock)
		if len(tmpBlock) != 0 {
			block2 = append(block2, tmpBlock)
		}
		tmpBlock = []rune{}

		for range fspace {
			tmpBlock = append(tmpBlock, '.')
		}
		block = slices.Concat(block, tmpBlock)
		if len(tmpBlock) != 0 {
			block2 = append(block2, tmpBlock)
		}
		tmpBlock = []rune{}

	}
	fmt.Println("part1: ", part1(block))
	fmt.Println("part2: ", part2(block2))
}

func part1(block []rune) int {
	lastIndex := len(block) - 1
	checksum := 0
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
	return checksum
}

func part2(block2 [][]rune) int {
	checksum := 0

	//print(block2)

	for tempLastIndex := len(block2) - 1; tempLastIndex >= 0; tempLastIndex-- {
		b := block2[tempLastIndex]
		if b[0] == '.' {
			continue
		}
		for i := 0; i < tempLastIndex; i++ {
			if block2[i][0] != '.' {
				continue
			}
			if len(block2[i]) >= len(b) {
				// Swap the blocks
				block2[tempLastIndex], block2[i] = block2[i], block2[tempLastIndex]
				if len(block2[tempLastIndex]) > len(b) {
					block2 = slices.Insert(block2, i+1, block2[tempLastIndex][:len(block2[tempLastIndex])-len(block2[i])])
					block2[tempLastIndex+1] = block2[tempLastIndex+1][:len(block2[i])]
					tempLastIndex++ // increment tempLastIndex to account for the new block added
				}
				break
			}
		}
		//print(block2)
	}

	i := 0
	for _, b := range block2 {
		if b[0] != '.' {
			for _, e := range b {
				checksum += i * (int(e) - '0')
				i++
			}
		} else {
			i += len(b)
		}
	}
	return checksum
}

func print(block2 [][]rune) {
	// Print the 2D slice in a formatted way for debugging purposes
	for _, x := range block2 {
		fmt.Printf("%s ", string(x))
	}
	fmt.Println()
}
