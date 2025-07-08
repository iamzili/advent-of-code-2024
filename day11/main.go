package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

var cache = make(map[int][]int)
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	stones := []int{}
	input, _ := os.ReadFile("input2")
	// Convert string to slice of ints
	json.Unmarshal([]byte("["+strings.ReplaceAll(string(input), " ", ",")+"]"), &stones)

	fmt.Println("part1: ", getCountAfterBlinking(stones, 25))
	fmt.Println("part2: ", getCountAfterBlinking(stones, 75))

}

func getCountAfterBlinking(stones []int, blinkCount int) int {
	count := 0
	for _, stone := range stones {
		count += getCountAfterXBlinking(stone, cache, blinkCount)
	}
	return count
}

func getCountAfterXBlinking(stone int, cache map[int][]int, blinkCount int) int {
	if _, ok := cache[stone]; ok {
		if cache[stone][blinkCount-1] != 0 {
			return cache[stone][blinkCount-1]
		}
	} else {
		cache[stone] = make([]int, 75)
	}

	if blinkCount == 1 {
		cache[stone][blinkCount-1] = len(changeStone(stone))
		return len(changeStone(stone))
	}

	sum := 0
	for _, stone := range changeStone(stone) {
		sum += getCountAfterXBlinking(stone, cache, blinkCount-1)
	}

	cache[stone][blinkCount-1] = sum
	return sum
}

func changeStone(stone int) []int {
	tmpStones := []int{}
	if stone == 0 {
		tmpStones = append(tmpStones, 1)
	} else if len(strconv.Itoa(stone))%2 == 0 {
		stoneString := strconv.Itoa(stone)
		i1, _ := strconv.Atoi(stoneString[:len(stoneString)/2])
		i2, _ := strconv.Atoi(stoneString[len(stoneString)/2:])
		tmpStones = append(tmpStones, i1, i2)
	} else {
		tmpStones = append(tmpStones, stone*2024)
	}
	return tmpStones
}
